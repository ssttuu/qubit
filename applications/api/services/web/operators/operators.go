package operators

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/operator"
	_ "github.com/stupschwartz/qubit/core/operators"
	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

var operatorsTable = "operators"

type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
	ComputeClient  compute_pb.ComputeClient
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client, computeClient compute_pb.ComputeClient) {
	operators_pb.RegisterOperatorsServer(grpcServer, &Server{
		ComputeClient:  computeClient,
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	newObject := operator.NewFromProto(in.Operator)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  operatorsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: operatorsTable,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	var obj operator.Operator
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: operatorsTable,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	var objectList operator.Operators
	err := apiutils.List(&apiutils.ListConfig{
		// Don't load parameters for list because it's a large column
		Columns: []string{"id", "scene_id", "context", "type", "name"},
		DB:      s.PostgresClient,
		Out:     &objectList,
		Table:   operatorsTable,
	})
	if err != nil {
		return nil, err
	}
	return &operators_pb.ListOperatorsResponse{Operators: objectList.ToProto()}, nil
}

func (s *Server) Render(ctx context.Context, in *operators_pb.RenderOperatorRequest) (*operators_pb.RenderOperatorResponse, error) {
	//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	//	operatorKey := datastore.NameKey(operator.Kind, in.Id, sceneKey)
	//
	//	// TODO: make gRPC request for the operator?
	//	var theOperator operator.Operator
	//	if err := s.DatastoreClient.Get(ctx, operatorKey, &theOperator); err != nil {
	//		return nil, errors.Wrapf(err, "Failed to get operator to be rendered, %v", operatorKey)
	//	}
	//
	//	listParamsRequest := &parameters_pb.ListParametersRequest{
	//		OrganizationId: in.OrganizationId,
	//		SceneId:        in.SceneId,
	//		OperatorId:     in.Id,
	//	}
	//
	//	params_pb, err := s.ParametersClient.List(ctx, listParamsRequest)
	//	if err != nil {
	//		return nil, errors.Wrapf(err, "Failed to list parameters, %v", listParamsRequest)
	//	}
	//
	//	renderImageRequest := &compute_pb.RenderImageRequest{
	//		Operator:   theOperator.ToProto(),
	//		Parameters: params_pb.Parameters,
	//	}
	//	renderImageResponse, err := s.ComputeClient.RenderImage(ctx, renderImageRequest)
	//	if err != nil {
	//		return nil, errors.Wrapf(err, "Failed to render operator, %v", theOperator)
	//	}
	//
	//	imagePlane := image.NewPlaneFromProto(renderImageResponse.ImagePlane)
	//
	//	//
	//	//// TODO: create bucket per Organization
	//	//// TODO: hash OrgId, SceneId, and OperatorId to get bucket path
	//	imageProtoObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.bytes", in.OrganizationId, in.SceneId, in.Id, in.Frame)
	//	imagePngObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.png", in.OrganizationId, in.SceneId, in.Id, in.Frame)
	//
	//	bucket := s.StorageClient.Bucket("qubit-dev-161916")
	//	imageProtoObject := bucket.Object(imageProtoObjectPath)
	//
	//	// PROTO
	//	protoWriter := imageProtoObject.NewWriter(ctx)
	//	protoWriter.ContentType = "application/octet-stream"
	//
	//	image_bytes, err := proto.Marshal(imagePlane.ToProto())
	//	if err != nil {
	//		return nil, errors.Wrapf(err, "Failed to marshal imagePlane proto, %v", imagePlane)
	//	}
	//
	//	protoWriter.Write(image_bytes)
	//	protoWriter.Close()
	//
	//	// PNG
	//	pngWriter := imageProtoObject.NewWriter(ctx)
	//	pngWriter.ContentType = "image/png"
	//
	//	if err := png.Encode(pngWriter, imagePlane.ToNRGBA()); err != nil {
	//		return nil, errors.Wrap(err, "Failed to encode png")
	//	}
	//
	//	pngWriter.Close()
	//
	//	return &operators_pb.RenderOperatorResponse{ResultUrl: imagePngObjectPath, ResultType: operator.IMAGE}, nil
	return nil, nil
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	newObject := operator.NewFromProto(in.Operator)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &operator.Operator{},
		Table:     operatorsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
