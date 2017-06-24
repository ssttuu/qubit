package image_sequences

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/image_sequence"
	image_sequences_pb "github.com/stupschwartz/qubit/proto-gen/go/image_sequences"
)

var image_sequencesTable = "image_sequences"

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	image_sequences_pb.RegisterImageSequencesServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func (s *Server) Create(ctx context.Context, in *image_sequences_pb.CreateImageSequenceRequest) (*image_sequences_pb.ImageSequence, error) {
	newObject := image_sequence.NewFromProto(in.ImageSequence)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  image_sequencesTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *image_sequences_pb.DeleteImageSequenceRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: image_sequencesTable,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *image_sequences_pb.GetImageSequenceRequest) (*image_sequences_pb.ImageSequence, error) {
	var obj image_sequence.ImageSequence
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: image_sequencesTable,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *image_sequences_pb.ListImageSequencesRequest) (*image_sequences_pb.ListImageSequencesResponse, error) {
	var objectList image_sequence.ImageSequences
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: image_sequencesTable,
	})
	if err != nil {
		return nil, err
	}
	return &image_sequences_pb.ListImageSequencesResponse{ImageSequences: objectList.ToProto()}, nil
}

func (s *Server) Update(ctx context.Context, in *image_sequences_pb.UpdateImageSequenceRequest) (*image_sequences_pb.ImageSequence, error) {
	newObject := image_sequence.NewFromProto(in.ImageSequence)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &image_sequence.ImageSequence{},
		Table:     image_sequencesTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
