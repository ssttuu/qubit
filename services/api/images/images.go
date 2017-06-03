package images
//
//import (
//	"math/rand"
//	"time"
//
//	"golang.org/x/net/context"
//	"cloud.google.com/go/datastore"
//	"github.com/golang/protobuf/ptypes/empty"
//	"github.com/pkg/errors"
//	"google.golang.org/grpc"
//
//	"github.com/stupschwartz/qubit/core/image"
//	images_pb "github.com/stupschwartz/qubit/proto-gen/go/images"
//	"github.com/stupschwartz/qubit/core/organization"
//	"github.com/stupschwartz/qubit/core/scene"
//	"github.com/stupschwartz/qubit/core/operator"
//	"os"
//	"google.golang.org/api/option"
//	"net"
//	"google.golang.org/grpc/grpclog"
//	"log"
//	"cloud.google.com/go/storage"
//)
//
//var r *rand.Rand
//
//func init() {
//	r = rand.New(rand.NewSource(time.Now().UnixNano()))
//}
//
//
//type Server struct {
//	DatastoreClient *datastore.Client
//	StorageClient *storage.Client
//}
//
//func (s *Server) List(ctx context.Context, in *images_pb.ListImagesRequest) (*images_pb.ListImagesResponse, error) {
//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
//	operatorKey := datastore.NameKey(operator.Kind, in.SceneId, sceneKey)
//
//	var images image.Frames
//	_, err := s.DatastoreClient.GetAll(ctx, datastore.NewQuery(image.Kind).Ancestor(operatorKey), &images)
//	if err != nil {
//		return nil, errors.Wrap(err, "Could not get all")
//	}
//
//	return &images_pb.ListImageResponse{Image:images.ToProto(), NextPageToken:""}, nil
//}
//
//func (s *Server) Get(ctx context.Context, in *images_pb.GetParameterRequest) (*images_pb.Parameter, error) {
//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
//	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
//	imageKey := datastore.NameKey(image.Kind, in.ParameterId, operatorKey)
//
//	var existingParameter image.Parameter
//	if err := s.DatastoreClient.Get(ctx, imageKey, &existingParameter); err != nil {
//		return nil, errors.Wrap(err, "Could not get datastore entity")
//	}
//
//	return existingParameter.ToProto(), nil
//}
//
//func (s *Server) Create(ctx context.Context, in *images_pb.CreateParameterRequest) (*images_pb.Parameter, error) {
//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
//	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
//	imageKey := datastore.NameKey(image.Kind, in.Parameter.Id, operatorKey)
//
//	newParameter := image.NewParameterFromProto(in.Parameter)
//
//	if _, err := s.DatastoreClient.Put(ctx, imageKey, &newParameter); err != nil {
//		return nil, errors.Wrapf(err, "Failed to put operator %v", newParameter.Id)
//	}
//
//	return newParameter.ToProto(), nil
//}
//
//func (s *Server) Update(ctx context.Context, in *images_pb.UpdateParameterRequest) (*images_pb.Parameter, error) {
//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
//	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
//	imageKey := datastore.NameKey(image.Kind, in.ParameterId, operatorKey)
//
//	newParameter := image.NewParameterFromProto(in.Parameter)
//
//	_, err := s.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
//		var existingParameter image.Parameter
//		if err := tx.Get(imageKey, &existingParameter); err != nil {
//			return errors.Wrapf(err, "Failed to get operator in tx %v", existingParameter)
//		}
//
//		existingParameter.Id = newParameter.Id
//
//		_, err := tx.Put(imageKey, &existingParameter)
//		if err != nil {
//			return errors.Wrapf(err, "Failed to put operator in tx %v", existingParameter)
//		}
//
//		newParameter = &existingParameter
//
//		return nil
//	})
//
//	if err != nil {
//		return nil, errors.Wrap(err, "Failed to update operator")
//	}
//
//	return newParameter.ToProto(), nil
//}
//
//func (s *Server) Delete(ctx context.Context, in *images_pb.DeleteParameterRequest) (*empty.Empty, error) {
//	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
//	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
//	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
//	imageKey := datastore.NameKey(image.Kind, in.ParameterId, operatorKey)
//
//	if err := s.DatastoreClient.Delete(ctx, imageKey); err != nil {
//		return nil, errors.Wrapf(err, "Failed to delete operator by id: %v", in.OperatorId)
//	}
//
//	return &empty.Empty{}, nil
//}
//
//func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client, storageClient *storage.Client) {
//	images_pb.RegisterImagesServer(grpcServer, &Server{
//		DatastoreClient: datastoreClient,
//		StorageClient: storageClient,
//	})
//}
