package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "github.com/stupschwartz/qubit/protos"
	"github.com/stupschwartz/qubit/operator"
	"github.com/stupschwartz/qubit/image"
	"github.com/stupschwartz/qubit/node"
	"github.com/stupschwartz/qubit/params"
	"os"
	"encoding/json"
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"log"
	"time"
)

type ComputeServer struct {
	DatastoreClient *datastore.Client
	StorageClient   *storage.Client
}

func (cs *ComputeServer) Render(ctx context.Context, renderRequest *pb.RenderRequest) (*pb.RenderResponse, error) {
	nodeUuid := renderRequest.GetNode().GetId()
	boundingBox := renderRequest.GetBoundingBox()
	nodeKey := datastore.NameKey("Node", nodeUuid, nil)

	var theNode node.Node
	if err := cs.DatastoreClient.Get(ctx, nodeKey, &theNode); err != nil {
		log.Fatalf("Failed to get node to be rendered, %v", err)
	}

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

	op := operator.GetOperation(theNode.Type)

	inputImagePlanes := make([]image.Plane, len(theNode.Inputs))

	for index, inputNodeId := range theNode.Inputs {
		inputImagePlanes[index] = MakeRenderRequest(ctx, inputNodeId, boundingBox)
	}

	imagePlane := op(inputImagePlanes, theParams, int(boundingBox.GetEndX()), int(boundingBox.GetEndY()))

	response := &pb.RenderResponse{ImagePlane: imagePlane.ToProto()}

	return response, nil
}

func MakeRenderRequest(ctx context.Context, nodeUuid string, bbox *pb.BoundingBox) image.Plane {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	// TODO: client side load balancing
	conn, err := grpc.Dial("compute:10000", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	computeClient := pb.NewComputeClient(conn)

	renderRequest := &pb.RenderRequest{Node: &pb.Node{Id: nodeUuid}, BoundingBox: bbox}

	renderResponse, err := computeClient.Render(ctx, renderRequest)
	if err != nil {
		grpclog.Fatal("failed to render")
	}

	return image.NewPlaneFromProto(renderResponse.GetImagePlane())
}

func newServer() *ComputeServer {
	s := new(ComputeServer)

	projID := os.Getenv("DATASTORE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "DATASTORE_PROJECT_ID"`)
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

	return s
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	//creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}
	//opts = []grpc.ServerOption{grpc.Creds(creds)}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterComputeServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
