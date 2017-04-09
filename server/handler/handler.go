package handler

import (
	"github.com/stupschwartz/qubit/server/env"
	"log"
	"net/http"
	"fmt"
	"golang.org/x/net/context"
	"cloud.google.com/go/trace"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Handler struct {
	*env.Env
	H func(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	span := h.Env.TraceClient.SpanFromRequest(r)
	defer span.Finish()

	ctx := trace.NewContext(context.Background(), span)

	err := h.H(ctx, h.Env, w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any env types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
