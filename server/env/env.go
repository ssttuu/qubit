package env

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	"github.com/stupschwartz/qubit/server/protos/parameters"
	"github.com/stupschwartz/qubit/server/protos/organizations"
	"github.com/stupschwartz/qubit/server/protos/scenes"
	"github.com/stupschwartz/qubit/server/protos/operators"
)

type Env struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
	TraceClient *trace.Client

	OrganizationsClient *organizations.OrganizationsClient
	ScenesClient *scenes.ScenesClient
	OperatorsClient *operators.OperatorsClient
	ParametersClient *parameters.ParametersClient
}
