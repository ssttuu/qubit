package images

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/image"
	images_pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

var imagesTable = "images"

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	images_pb.RegisterImagesServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func (s *Server) Create(ctx context.Context, in *images_pb.CreateImageRequest) (*images_pb.Image, error) {
	newObject := image.NewFromProto(in.Image)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  imagesTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *images_pb.DeleteImageRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: imagesTable,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *images_pb.GetImageRequest) (*images_pb.Image, error) {
	var obj image.Image
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: imagesTable,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *images_pb.ListImagesRequest) (*images_pb.ListImagesResponse, error) {
	var objectList image.Images
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: imagesTable,
	})
	if err != nil {
		return nil, err
	}
	return &images_pb.ListImagesResponse{Images: objectList.ToProto()}, nil
}

func (s *Server) Update(ctx context.Context, in *images_pb.UpdateImageRequest) (*images_pb.Image, error) {
	newObject := image.NewFromProto(in.Image)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &image.Image{},
		Table:     imagesTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
