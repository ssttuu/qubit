package main

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/params"
	pb "github.com/stupschwartz/qubit/compute/protos/compute"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"fmt"
)

type ComputeServer struct {
	DatastoreClient *datastore.Client
	StorageClient   *storage.Client
	TraceClient     *trace.Client
	ComputeClient   *pb.ComputeClient
}

func (cs *ComputeServer) Render(stream pb.Compute_RenderServer) error {
	ctx := stream.Context()

	grp := sync.WaitGroup{}
	mu := sync.Mutex{}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		grp.Add(1)
		go func(ctx context.Context, sceneUuid string, operatorUuid string, boundingBox *pb.BoundingBox, stream pb.Compute_RenderServer, mu *sync.Mutex) {
			cs.ActuallyRender(ctx, sceneUuid, operatorUuid, boundingBox, stream, mu)
			grp.Done()
		}(ctx, in.GetSceneId(), in.GetOperatorId(), in.GetBoundingBox(), stream, &mu)
	}

	grp.Wait()

	return nil
}

func (cs *ComputeServer) ActuallyRender(ctx context.Context, sceneUuid string, operatorUuid string, boundingBox *pb.BoundingBox, stream pb.Compute_RenderServer, mu *sync.Mutex) error {
	sceneKey := datastore.NameKey("Scene", sceneUuid, nil)
	operatorKey := datastore.NameKey("Operator", operatorUuid, sceneKey)

	var theOperator operator.Operator
	if err := cs.DatastoreClient.Get(ctx, operatorKey, &theOperator); err != nil {
		log.Fatalf("Failed to get operator to be rendered, %v", err)
	}

	var theParams params.Parameters
	bucket := cs.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))
	paramsObj := bucket.Object(fmt.Sprintf("params/%s", theOperator.Id))

	reader, err := paramsObj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create Reader, %v", err)
	}

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&theParams); err != nil {
		log.Fatal(err)
	}
	reader.Close()

	//renderInputsSpan := span.NewChild("render:" + stringBBox)
	//inputImagePlanes := make([]image.Plane, len(theOperator.Inputs))
	//for index, inputOperatorId := range theOperator.Inputs {
	//	inputImagePlanes[index] = cs.RenderInput(ctx, sceneUuid, inputOperatorId, boundingBox)
	//}
	//renderInputsSpan.Finish()

	op := operator.GetOperation(theOperator.Type)
	imagePlane := op([]image.Plane{}, theParams, boundingBox.GetStartX(), boundingBox.GetStartY(), boundingBox.GetEndX(), boundingBox.GetEndY())

	response := &pb.RenderResponse{ImagePlane: imagePlane.ToProto(), BoundingBox: boundingBox}

	//fmt.Printf("Sending Response: %v\n", boundingBox)
	//fmt.Printf(" -- w: %v, h: %v, cs: %v, rs: %v, data: %v\n", imagePlane.Width, imagePlane.Height, len(imagePlane.Components), len(imagePlane.Components[0].Rows), len(imagePlane.Components[0].Rows[0].Data))

	mu.Lock()
	err = stream.Send(response)
	mu.Unlock()

	if err != nil {
		return errors.Wrap(err, "Failed to send stream")
	}

	return nil
}

//func (cs *ComputeServer) RenderInput(ctx context.Context, sceneUuid string, operatorUuid string, bbox *pb.BoundingBox) image.Plane {
//	var opts []grpc.DialOption
//	opts = append(opts, grpc.WithInsecure())
//	opts = append(opts, grpc.WithUnaryInterceptor(trace.GRPCClientInterceptor()))
//
//	// TODO: client side load balancing
//	conn, err := grpc.Dial("compute:10000", opts...)
//	if err != nil {
//		grpclog.Fatalf("fail to dial: %v", err)
//	}
//	defer conn.Close()
//	computeClient := pb.NewComputeClient(conn)
//
//	renderResponse, err := computeClient.Render(ctx, &pb.RenderRequest{SceneId: sceneUuid, OperatorId: operatorUuid, BoundingBox: bbox})
//	if err != nil {
//		grpclog.Fatal("failed to render")
//	}
//
//	return image.NewPlaneFromProto(renderResponse.GetImagePlane())
//}

func newServer(traceClient *trace.Client) *ComputeServer {
	s := new(ComputeServer)

	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
	}

	ctx := context.Background()

	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	datastoreClient, err := datastore.NewClient(ctx, projID, serviceCredentials)
	for err != nil {
		log.Printf("Could not create datastore client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		datastoreClient, err = datastore.NewClient(ctx, projID, serviceCredentials)
	}

	storageClient, err := storage.NewClient(ctx, serviceCredentials)
	for err != nil {
		log.Printf("Could not create storage client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		storageClient, err = storage.NewClient(ctx, serviceCredentials)
	}

	s.DatastoreClient = datastoreClient
	s.StorageClient = storageClient
	s.TraceClient = traceClient

	log.Println("NewServer")

	return s
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()

	traceClient, err := trace.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"))
	if err != nil {
		log.Fatalf("Could not create trace client: %v\n", err)
	}
	p, err := trace.NewLimitedSampler(1.0, 10)
	if err != nil {
		log.Fatalf("Could not create limited sampler: %v\n", err)
	}
	traceClient.SetSamplingPolicy(p)

	log.Println("MAIN")

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(serverInterceptor(traceClient)))
	pb.RegisterComputeServer(grpcServer, newServer(traceClient))
	grpcServer.Serve(lis)
}

func serverInterceptor(traceClient *trace.Client) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			md = metadata.New(nil)
		}
		header := strings.Join(md["x-cloud-trace-context"], "")

		span := traceClient.SpanFromHeader(info.FullMethod, header)
		defer span.Finish()

		return handler(srv, ss)
	}
}
