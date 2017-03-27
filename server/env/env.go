package env

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
)

type Env struct {
	DatastoreClient *datastore.Client
	StorageClient *storage.Client
}
