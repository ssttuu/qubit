package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bmizerany/pq"
	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc/status"

	"github.com/stupschwartz/qubit/core/computation"
	"github.com/stupschwartz/qubit/core/computation_status"
	computations_pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

const pubSubCoordinatorSubscriptionID = "coordinator"

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
	topic, err := pubSubClient.CreateTopic(ctx, computation.PubSubTopicID)
	if err != nil {
		// 409 ALREADY_EXISTS is an inevitable and harmless error
		// https://cloud.google.com/pubsub/docs/reference/error-codes
		if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != 409 {
			return nil, errors.Wrapf(err, "Failed to get-or-create topic %v", computation.PubSubTopicID)
		}
	}
	// Default of 10 second ack deadline
	subscription, err := pubSubClient.CreateSubscription(ctx, pubSubCoordinatorSubscriptionID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		// 409 ALREADY_EXISTS is an inevitable and harmless error
		// https://cloud.google.com/pubsub/docs/reference/error-codes
		if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != 409 {
			return nil, errors.Wrapf(err, "Failed to get-or-create subscription %v", pubSubCoordinatorSubscriptionID)
		}
	}
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var msgPBCompStatus computations_pb.ComputationStatus
		err := proto.Unmarshal(msg.Data, &msgPBCompStatus)
		if err != nil {
			log.Println(err)
			msg.Ack()
			return
		}
		msgCompStatus := computation_status.NewFromProto(msgPBCompStatus)
		// TODO: Where to ack?
		log.Println("ACK:", msgCompStatus)
		msg.Ack()
		// TODO: Use transaction
		var compStatus computation_status.ComputationStatus
		err = postgresClient.Get(&compStatus, fmt.Sprintf("SELECT * FROM %v WHERE id=$1", computation_status.TableName), msgCompStatus.Id)
		if err != nil {
			log.Printf("could not get computation status with ID: %v", msgCompStatus.Id)
			return
		}
		var comp computation.Computation
		err = postgresClient.Get(&comp, fmt.Sprintf("SELECT * FROM %v WHERE id=$1", computation.TableName), compStatus.ComputationId)
		if err != nil {
			log.Printf("could not get computation with ID: %v", compStatus.ComputationId)
			return
		}
		log.Println("COMPUTATION:", comp)
		// TODO: Business logic
	})
	if err != nil {
		log.Fatal(err)
	}
}
