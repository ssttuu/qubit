package operators

import (
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	_ "github.com/stupschwartz/qubit/core/operators"
	operators_pb "github.com/stupschwartz/qubit/server/protos/operators"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"fmt"
	"github.com/stupschwartz/qubit/core/params"
	"os"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/golang/protobuf/proto"
	"image/png"
)

const OrganizationKind string = "Organization"
const SceneKind string = "Scene"
const OperatorKind string = "Operator"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)

	var operators operator.Operators
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery(OperatorKind).Ancestor(sceneKey), &operators)
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
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)

	var existingOperator operator.Operator
	if err := s.env.DatastoreClient.Get(ctx, operatorKey, &existingOperator); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingOperator.ToProto()
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	in.Operator.Id = fmt.Sprint(r.Int63())
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.Operator.Id, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	if _, err := s.env.DatastoreClient.Put(ctx, operatorKey, &newOperator); err != nil {
		return nil, errors.Wrapf(err, "Failed to put operator %v", newOperator.Id)
	}

	return newOperator.ToProto()
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
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
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)

	if err := s.env.DatastoreClient.Delete(ctx, operatorKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to delete operator by id: %v", in.OperatorId)
	}

	return &empty.Empty{}, nil
}

func (s *Server) Render(ctx context.Context, in *operators_pb.RenderOperatorRequest) (*operators_pb.RenderOperatorResponse, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)

	// TODO: make gRPC request for the operator?
	var theOperator operator.Operator
	if err := s.env.DatastoreClient.Get(ctx, operatorKey, &theOperator); err != nil {
		return nil, errors.Wrapf(err, "Failed to get operator to be rendered, %v", operatorKey)
	}

	// TODO: make gRPC request for the parameters?
	var theParams params.Parameters = params.Parameters{}
	bucket := s.env.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))

	// TODO: create bucket per Organization
	// TODO: hash OrgId, SceneId, and OperatorId to get bucket path
	//paramsObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/params.json", in.OrganizationId, in.SceneId, in.OperatorId)
	//paramsObj := bucket.Object(paramsObjectPath)
	//
	//reader, err := paramsObj.NewReader(ctx)
	//if err != nil {
	//	return nil, errors.Wrapf(err, "Failed to create Reader, %v", paramsObjectPath)
	//}
	//
	//decoder := json.NewDecoder(reader)
	//if err := decoder.Decode(&theParams); err != nil {
	//	return nil, errors.Wrapf(err, "Failed to decode params, %v", paramsObjectPath)
	//}
	//reader.Close()

	op, err := operator.GetOperation(theOperator.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Operation for rendering, %v", theOperator.Type)
	}
	imagePlane := op.Process([]image.Plane{}, theParams, in.BoundingBox.StartX, in.BoundingBox.StartY, in.BoundingBox.EndX, in.BoundingBox.EndY)

	// TODO: create bucket per Organization
	// TODO: hash OrgId, SceneId, and OperatorId to get bucket path
	imageProtoObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.bytes", in.OrganizationId, in.SceneId, in.OperatorId, in.Frame)
	imagePngObjectPath := fmt.Sprintf("organizations/%d/scenes/%d/operators/%d/images/%d/image.png", in.OrganizationId, in.SceneId, in.OperatorId, in.Frame)

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

func newServer(e *env.Env) *Server {
	return &Server{
		env: e,
	}
}

func Register(server *grpc.Server, e *env.Env) {
	operators_pb.RegisterOperatorsServer(server, newServer(e))
}
