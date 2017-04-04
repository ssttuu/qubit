package env

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	pb "github.com/stupschwartz/qubit/protos"
)

type Env struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
	ComputeClient pb.ComputeClient
}
