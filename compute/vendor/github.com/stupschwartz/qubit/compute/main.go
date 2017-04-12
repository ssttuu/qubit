package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "github.com/stupschwartz/qubit/protos"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/node"
	"github.com/stupschwartz/qubit/core/params"
	"os"
	"encoding/json"
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	"google.golang.org/api/option"
	"log"
	"time"
)

type ComputeServer struct {
	DatastoreClient *datastore.Client
	StorageClient   *storage.Client
	ComputeClient   *pb.ComputeClient
}

func (cs *ComputeServer) Render(ctx context.Context, renderRequest *pb.RenderRequest) (*pb.RenderResponse, error) {
	sceneUuid := renderRequest.GetSceneId()
	nodeUuid := renderRequest.GetNodeId()
	boundingBox := renderRequest.GetBoundingBox()

	sceneKey := datastore.NameKey("Scene", sceneUuid, nil)
	nodeKey := datastore.NameKey("Node", nodeUuid, sceneKey)

	span := trace.FromContext(ctx).NewChild("compute.Render")
	defer span.Finish()

	datastoreGet := span.NewChild("datastore.Get")
	var theNode node.Node
	if err := cs.DatastoreClient.Get(ctx, nodeKey, &theNode); err != nil {
		log.Fatalf("Failed to get node to be rendered, %v", err)
	}
	datastoreGet.Finish()

	storageGet := span.NewChild("storage.Get")
	var theParams params.Parameters
	bucket := cs.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))
	paramsObj := bucket.Object("params/" + theNode.Id)

	reader, err := paramsObj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create Reader, %v", err)
	}

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&theParams); err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	storageGet.Finish()

	op := operator.GetOperation(theNode.Type)

	inputImagePlanes := make([]image.Plane, len(theNode.Inputs))

	for index, inputNodeId := range theNode.Inputs {
		inputImagePlanes[index] = cs.RenderInput(ctx, sceneUuid, inputNodeId, boundingBox)
	}

	renderOp := span.NewChild("op.Render")
	imagePlane := op(inputImagePlanes, theParams, boundingBox.GetEndX(), boundingBox.GetEndY())
	renderOp.Finish()

	serialize := span.NewChild("serialize.response")
	response := &pb.RenderResponse{ImagePlane: imagePlane.ToProto()}
	serialize.Finish()

	return response, nil
}

func (cs *ComputeServer) RenderInput(ctx context.Context, sceneUuid string, nodeUuid string, bbox *pb.BoundingBox) image.Plane {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(trace.GRPCClientInterceptor()))

	// TODO: client side load balancing
	conn, err := grpc.Dial("compute:10000", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	computeClient := pb.NewComputeClient(conn)

	renderResponse, err := computeClient.Render(ctx, &pb.RenderRequest{SceneId: sceneUuid, NodeId: nodeUuid, BoundingBox: bbox})
	if err != nil {
		grpclog.Fatal("failed to render")
	}

	return image.NewPlaneFromProto(renderResponse.GetImagePlane())
}

func newServer() *ComputeServer {
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

	log.Println("NewServer")

	return s
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
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

	//var opts []grpc.ServerOption
	//opts = append(opts, EnableGRPCTracingServerOption(traceClient))
	//creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}
	//opts = []grpc.ServerOption{grpc.Creds(creds)}

	log.Println("MAIN")

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(trace.GRPCServerInterceptor(traceClient)))
	pb.RegisterComputeServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
