package parameters

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/parameter"
	parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *parameters_pb.ListParametersRequest) (*parameters_pb.ListParametersResponse, error) {
	// TODO: Permissions
	var params parameter.Parameters
	err := s.PostgresClient.Select(&params, "SELECT * FROM parameters")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select parameters")
	}
	return &parameters_pb.ListParametersResponse{Parameters: params.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *parameters_pb.GetParameterRequest) (*parameters_pb.Parameter, error) {
	// TODO: Permissions
	var param parameter.Parameter
	err := s.PostgresClient.Get(&param, "SELECT * FROM parameters WHERE id=$1", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get parameter with ID %v", in.Id)
	}
	return param.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *parameters_pb.CreateParameterRequest) (*parameters_pb.Parameter, error) {
	// TODO: Validation
	result, err := s.PostgresClient.NamedExec(
		`INSERT INTO parameters () VALUES ()`,
		map[string]interface{}{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create parameter")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve new ID")
	}
	newParameter := parameter.Parameter{
		Id: id,
	}
	return newParameter.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *parameters_pb.UpdateParameterRequest) (*parameters_pb.Parameter, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for parameter with ID %v", in.Id)
	}
	txStmt, err := tx.Prepare(`SELECT * FROM parameters WHERE id=? FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select parameter in tx %v", in.Id)
	}
	row := txStmt.QueryRow(in.Id)
	if row == nil {
		return nil, errors.Wrapf(err, "No parameter with ID %v exists", in.Id)
	}
	var existingParameter parameter.Parameter
	err = row.Scan(&existingParameter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load parameter from row")
	}
	// TODO: Make update fields dynamic
	// TODO: Is there anything to update?
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update parameter")
	}
	return existingParameter.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *parameters_pb.DeleteParameterRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	_, err := s.PostgresClient.Queryx("DELETE FROM parameters WHERE id=?", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted parameter by id: %v", in.Id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	parameters_pb.RegisterParametersServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func NewClient(conn *grpc.ClientConn) parameters_pb.ParametersClient {
	return parameters_pb.NewParametersClient(conn)
}
