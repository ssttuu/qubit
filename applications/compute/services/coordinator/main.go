package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/api/option"

	"github.com/stupschwartz/qubit/core/computation"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal(`You need to set the environment variable "POSTGRES_URL"`)
	}
	parsedURL, err := pq.ParseURL(postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	postgresClient, err := sqlx.Open("postgres", parsedURL)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	pubSubClient, err := pubsub.NewClient(ctx, projID, serviceCredentials)
	for err != nil {
		log.Printf("Could not create pubsub client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		pubSubClient, err = pubsub.NewClient(ctx, projID, serviceCredentials)
	}
	topic, _ := pubSubClient.CreateTopic(ctx, computation.PubSubTopicID)
	subscriptionID := "coordinator"
	// Default of 10 second ack deadline
	subscription, _ := pubSubClient.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var messageData map[string]string
		// TODO: Use gRPC for serialization of messages instead of JSON
		if err := json.Unmarshal(msg.Data, &messageData); err != nil {
			log.Printf("could not decode message data: %#v", msg)
			msg.Ack()
			return
		}
		log.Println("ACK:", messageData)
		var comp computation.Computation
		err = postgresClient.Get(&comp, fmt.Sprintf("SELECT * FROM %v WHERE id=$1", computation.TableName), messageData["computation_id"])
		if err != nil {
			log.Printf("could not get computation with ID: %v", messageData["computation_id"])
			msg.Ack()
			return
		}
		log.Println("COMPUTATION:", comp)
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
}
