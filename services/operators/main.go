package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"image/png"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/api/option"
	"github.com/golang/protobuf/proto"
	"cloud.google.com/go/storage"

	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
	_ "github.com/stupschwartz/qubit/core/operators"
	"github.com/stupschwartz/qubit/core/parameter"
	"github.com/stupschwartz/qubit/core/organization"
	"github.com/stupschwartz/qubit/core/scene"
	parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
)


var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
	ParametersClient parameters_pb.ParametersClient
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)

	var operators operator.Operators
	_, err := s.DatastoreClient.GetAll(ctx, datastore.NewQuery(operator.Kind).Ancestor(sceneKey), &operators)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	operators_proto, err := operators.ToProto()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert operators to proto, %v", operators_proto)
	}

	return &operators_pb.ListOperatorsResponse{Operators:operators_proto, NextPageToken:""}, nil
}

func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)

	var existingOperator operator.Operator
	if err := s.DatastoreClient.Get(ctx, operatorKey, &existingOperator); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingOperator.ToProto()
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	in.Operator.Id = fmt.Sprint(r.Int63())
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.Operator.Id, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	if _, err := s.DatastoreClient.Put(ctx, operatorKey, &newOperator); err != nil {
		return nil, errors.Wrapf(err, "Failed to put operator %v", newOperator.Id)
	}

	return newOperator.ToProto()
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	_, err := s.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingOperator operator.Operator
		if err := tx.Get(operatorKey, &existingOperator); err != nil {
			return errors.Wrapf(err, "Failed to get operator in tx %v", existingOperator)
		}

		existingOperator.Name = newOperator.Name

		_, err := tx.Put(operatorKey, &existingOperator)
		if err != nil {
			return errors.Wrapf(err, "Failed to put operator in tx %v", existingOperator)
		}

		newOperator = existingOperator

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update operator")
	}

	return newOperator.ToProto()
}

func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)

	if err := s.DatastoreClient.Delete(ctx, operatorKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to delete operator by id: %v", in.OperatorId)
	}

	return &empty.Empty{}, nil
}

func (s *Server) Render(ctx context.Context, in *operators_pb.RenderOperatorRequest) (*operators_pb.RenderOperatorResponse, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)

	// TODO: make gRPC request for the operator?
	var theOperator operator.Operator
	if err := s.DatastoreClient.Get(ctx, operatorKey, &theOperator); err != nil {
		return nil, errors.Wrapf(err, "Failed to get operator to be rendered, %v", operatorKey)
	}

	listParamsRequest := &parameters_pb.ListParametersRequest{
		OrganizationId: in.OrganizationId,
		SceneId: in.SceneId,
		OperatorId: in.OperatorId,
	}

	params_pb, err := s.ParametersClient.List(ctx, listParamsRequest)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to list parameters, %v", listParamsRequest)
	}

	params := parameter.NewParametersFromProto(params_pb.Parameters)



	op, err := operator.GetOperation(theOperator.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Operation for rendering, %v", theOperator.Type)
	}

	imagePlane, err := op.Process([]*image.Plane{}, params, in.BoundingBox.StartX, in.BoundingBox.StartY, in.BoundingBox.EndX, in.BoundingBox.EndY)
	//
	//// TODO: create bucket per Organization
	//// TODO: hash OrgId, SceneId, and OperatorId to get bucket path
	imageProtoObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.bytes", in.OrganizationId, in.SceneId, in.OperatorId, in.Frame)
	imagePngObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.png", in.OrganizationId, in.SceneId, in.OperatorId, in.Frame)

	bucket := s.StorageClient.Bucket("qubit-dev-161916")
	imageProtoObject := bucket.Object(imageProtoObjectPath)

	// PROTO
	protoWriter := imageProtoObject.NewWriter(ctx)
	protoWriter.ContentType = "application/octet-stream"

	image_bytes, err := proto.Marshal(imagePlane.ToProto())
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to marshal imagePlane proto, %v", imagePlane)
	}

	protoWriter.Write(image_bytes)
	protoWriter.Close()

	// PNG
	pngWriter := imageProtoObject.NewWriter(ctx)
	pngWriter.ContentType = "image/png"

	if err := png.Encode(pngWriter, imagePlane.ToNRGBA()); err != nil {
		return nil, errors.Wrap(err, "Failed to encode png")
	}

	pngWriter.Close()

	return &operators_pb.RenderOperatorResponse{ResultUrl:imagePngObjectPath, ResultType: operator.IMAGE }, nil
}


func main() {
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

	conn, err := grpc.Dial("parameters:9000", grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Parameters service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial("parameters:9000", grpc.WithInsecure())
	}
	defer conn.Close()

	parametersClient := parameters_pb.NewParametersClient(conn)


	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	operators_pb.RegisterOperatorsServer(grpcServer, &Server{
		DatastoreClient: datastoreClient,
		StorageClient: storageClient,
		ParametersClient: parametersClient,
	})

	grpcServer.Serve(lis)
}
