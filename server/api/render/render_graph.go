package render

import (
	"github.com/stupschwartz/qubit/node"
	"log"
	"cloud.google.com/go/datastore"
	"github.com/stupschwartz/qubit/server/env"
	"context"
)

type RenderTask struct {
	Node *node.Node

}

func RenderNodeAndDependents(e *env.Env, nodeUuid string) {
	nodeKey := datastore.NameKey("Node", nodeUuid, nil)

	var existingNode node.Node
	if err := e.DatastoreClient.Get(context.Background(), nodeKey, &existingNode); err != nil {
		log.Fatalf("Failed to get node to be rendered, %v", err)
	}

	//taskUuid := uuid.NewV4()
	//taskKey := datastore.NameKey("Task", taskUuid.String(), nil)



	log.Printf("Rendering Task %v", existingNode)
}
