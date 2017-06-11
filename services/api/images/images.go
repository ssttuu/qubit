package images

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/image"
	images_pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
}

func (s *Server) List(ctx context.Context, in *images_pb.ListImagesRequest) (*images_pb.ListImagesResponse, error) {
	// TODO: Permissions
	var imageList image.Images
	err := s.PostgresClient.Select(&imageList, "SELECT * FROM images")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select images")
	}
	return &images_pb.ListImagesResponse{Images: imageList.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *images_pb.GetImageRequest) (*images_pb.Image, error) {
	// TODO: Permissions
	var im image.Image
	err := s.PostgresClient.Get(&im, "SELECT * FROM images WHERE id=$1", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get image with ID %v", in.Id)
	}
	return im.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *images_pb.CreateImageRequest) (*images_pb.Image, error) {
	// TODO: Validation
	result, err := s.PostgresClient.NamedExec(
		`INSERT INTO images (image_sequence_id, name, width, height, labels, planes)
			VALUES (:image_sequence_id, :name, :width, :height, :labels, :planes)`,
		map[string]interface{}{
			"image_sequence_id": in.Image.ImageSequenceId,
			"name":              in.Image.Name,
			"width":             in.Image.Width,
			"height":            in.Image.Height,
			"labels":            in.Image.Labels,
			"planes":            in.Image.Planes,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create image, %s", in.Image.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve new ID")
	}
	newImage := image.Image{
		Id:   string(id),
		Name: in.Image.Name,
	}
	return newImage.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *images_pb.UpdateImageRequest) (*images_pb.Image, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for image with ID %v", in.Id)
	}
	txStmt, err := tx.Prepare(`SELECT * FROM images WHERE id=? FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select image in tx %v", in.Id)
	}
	row := txStmt.QueryRow(in.Id)
	if row == nil {
		return nil, errors.Wrapf(err, "No image with ID %v exists", in.Id)
	}
	var existingImage image.Image
	err = row.Scan(&existingImage)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load image from row")
	}
	// TODO: Make update fields dynamic
	newImage := image.NewImageFromProto(in.Image)
	if newImage.Name != existingImage.Name {
		existingImage.Name = newImage.Name
		_, err = tx.Exec(
			`UPDATE images SET name=?, width=?, height=?, labels=?, planes=? WHERE id=?`,
			newImage.Name,
			newImage.Width,
			newImage.Height,
			newImage.Labels,
			newImage.Planes,
			in.Id,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update image with ID %v", in.Id)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update image")
	}
	return existingImage.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *images_pb.DeleteImageRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	_, err := s.PostgresClient.Queryx("DELETE FROM images WHERE id=?", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted image by id: %v", in.Id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client) {
	images_pb.RegisterImagesServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}
