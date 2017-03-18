package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitly/go-nsq"
)

// NoopNSQLogger allows us to pipe NSQ logs to dev/null
// The default NSQ logger is great for debugging, but did
// not fit our normally well structured JSON logs. Luckily
// NSQ provides a simple interface for injecting your own
// logger.
type NoopNSQLogger struct{}

// Output allows us to implement the nsq.Logger interface
func (l *NoopNSQLogger) Output(calldepth int, s string) error {
	return nil
}

// MessageHandler adheres to the nsq.Handler interface.
// This allows us to define our own custome handlers for
// our messages. Think of these handlers much like you would
// an http handler.
type MessageHandler struct{}

// HandleMessage is the only requirement needed to fulfill the
// nsq.Handler interface. This where you'll write your message
// handling logic.
func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// returning an error results in the message being re-enqueued
		// a REQ is sent to nsqd
		return errors.New("body is blank re-enqueue message")
	}

	// Let's log our message!
	log.Println(string(m.Body))

	// Returning nil signals to the consumer that the message has
	// been handled with success. A FIN is sent to nsqd
	return nil
}

func main() {
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer("node_changed", "compute", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.ChangeMaxInFlight(1)

	consumer.SetLogger(
		&NoopNSQLogger{},
		nsq.LogLevelError,
	)

	consumer.AddConcurrentHandlers(&MessageHandler{}, 1)

	if err := consumer.ConnectToNSQLookupd("nsqlookupd:4161"); err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	for {
		select {
		case <-consumer.StopChan:
			return
		case <-shutdown:
			consumer.Stop()
		}
	}
}