package env

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	"github.com/stupschwartz/qubit/compute/protos/compute"
)

type Env struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
	TraceClient *trace.Client
	ComputeClient compute.ComputeClient
}
