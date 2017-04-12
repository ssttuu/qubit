package env

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	pb "github.com/stupschwartz/qubit/protos"
)

type Env struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
	TraceClient *trace.Client
	ComputeClient pb.ComputeClient
}
