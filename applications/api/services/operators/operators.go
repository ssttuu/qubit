package operators

import (
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/operator"
	_ "github.com/stupschwartz/qubit/core/operators"
	"github.com/stupschwartz/qubit/core/parameter"
	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

// Server implements `service Health`.
type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
	ComputeClient  compute_pb.ComputeClient
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	// TODO: Permissions
	var ops operator.Operators
	err := s.PostgresClient.Select(&ops, "SELECT id, scene_id, context, type, name FROM operators")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select operators")
	}
	return &operators_pb.ListOperatorsResponse{Operators: ops.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	// TODO: Permissions
	operatorId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	var op operator.Operator
	err = s.PostgresClient.Get(&op, "SELECT id, scene_id, context, type, name FROM operators WHERE id=$1", operatorId)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get operator with ID %v", operatorId)
	}
	return op.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	// TODO: Validation
	query := `INSERT INTO operators (scene_id, context, type, name)
			  VALUES (:scene_id, :context, :type, :name)
			  RETURNING id`
	stmt, err := s.PostgresClient.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	var id int64
	err = stmt.Get(&id, map[string]interface{}{
		"scene_id": in.Operator.SceneId,
		"context":  in.Operator.Context,
		"type":     in.Operator.Type,
		"name":     in.Operator.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create operator, %s", in.Operator.Name)
	}
	newOp := operator.Operator{
		Id:         strconv.FormatInt(id, 10),
		Context:    in.Operator.Context,
		Type:       in.Operator.Type,
		Name:       in.Operator.Name,
		Parameters: parameter.NewParametersFromProto(in.Operator.Parameters),
	}
	return newOp.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	// TODO: Permissions & validation
	operatorId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for operator with ID %v", operatorId)
	}
	txStmt, err := tx.Prepare(`SELECT id, context, type, name FROM operators WHERE id=$1 FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select operator in tx %v", operatorId)
	}
	row := txStmt.QueryRow(operatorId)
	if row == nil {
		return nil, errors.Wrapf(err, "No operator with ID %v exists", operatorId)
	}
	var existingOperator operator.Operator
	// TODO: Scan parameters
	err = row.Scan(&existingOperator.Id, &existingOperator.Context, &existingOperator.Type, &existingOperator.Name)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load operator from row")
	}
	// TODO: Make update fields dynamic
	newOperator := operator.NewFromProto(in.Operator)
	// TODO: Handle parameters
	if newOperator.Context != existingOperator.Context || newOperator.Type != existingOperator.Type || newOperator.Name != existingOperator.Name {
		existingOperator.Context = newOperator.Context
		existingOperator.Type = newOperator.Type
		existingOperator.Name = newOperator.Name
		_, err = tx.Exec(
			`UPDATE operators SET context=$1, type=$2, name=$3 WHERE id=$4`,
			newOperator.Context,
			newOperator.Type,
			newOperator.Name,
			operatorId,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update operator with ID %v", operatorId)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update operator")
	}
	return existingOperator.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	operatorId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	_, err = s.PostgresClient.Queryx("DELETE FROM operators WHERE id=$1", operatorId)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted operator by id: %v", operatorId)
	}
	return &empty.Empty{}, nil
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

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client, computeClient compute_pb.ComputeClient) {
	operators_pb.RegisterOperatorsServer(grpcServer, &Server{
		ComputeClient:  computeClient,
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}
