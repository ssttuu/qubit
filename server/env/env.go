package env

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type Env struct {
	DatastoreClient *datastore.Client
	Context context.Context
}
