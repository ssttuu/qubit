package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"

	"context"
	"github.com/stupschwartz/qubit/image"
	"image/png"
	pb "github.com/stupschwartz/qubit/protos"
	"strconv"
	"github.com/pkg/errors"
)

func PostHandler(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	queryParams := r.URL.Query()

	nodeUuid := vars["id"]

	width, err := strconv.ParseInt(queryParams.Get("width"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "Failed to parse width")
	}

	height, err := strconv.ParseInt(queryParams.Get("height"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "Failed to parse height")
	}

	renderBounds := &pb.BoundingBox{StartX: 0, StartY: 0, EndX: width, EndY: height}

	renderRequest := &pb.RenderRequest{Node: &pb.Node{Id: nodeUuid}, BoundingBox: renderBounds}

	renderResponse, err := e.ComputeClient.Render(context.Background(), renderRequest)
	if err != nil {
		return errors.Wrap(err, "Failed to render")
	}

	imagePlane := image.NewPlaneFromProto(renderResponse.GetImagePlane())

	w.Header().Set("Content-Type", "image/png")

	if err := png.Encode(w, imagePlane.ToNRGBA()); err != nil {
		return errors.Wrap(err, "Failed to encode png")
	}

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/render").Subrouter()

	s.Handle("/{id}/", handler.Handler{environ, PostHandler}).Methods("POST")
}
