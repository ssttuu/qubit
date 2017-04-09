package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"github.com/stupschwartz/qubit/core/image"
	pb "github.com/stupschwartz/qubit/protos"
	"net/http"

	"golang.org/x/net/context"
	"image/png"
	"strconv"
	"github.com/pkg/errors"
	"cloud.google.com/go/trace"
)


func PostHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PostHandler")
	defer span.Finish()

	vars := mux.Vars(r)
	queryParams := r.URL.Query()

	sceneId := vars["scene_id"]
	nodeId := vars["node_id"]

	width, err := strconv.ParseInt(queryParams.Get("width"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "Failed to parse width")
	}

	height, err := strconv.ParseInt(queryParams.Get("height"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "Failed to parse height")
	}

	serializeSpan := span.NewChild("Serialize gRPC Request")
	renderBounds := &pb.BoundingBox{StartX: 0, StartY: 0, EndX: width, EndY: height}
	renderRequest := &pb.RenderRequest{Scene: &pb.Scene{Id: sceneId}, Node: &pb.Node{Id: nodeId}, BoundingBox: renderBounds}
	serializeSpan.Finish()

	renderResponse, err := e.ComputeClient.Render(ctx, renderRequest)
	if err != nil {
		return errors.Wrap(err, "Failed to render")
	}

	deserializeSpan := span.NewChild("Deserialize gRPC Request")
	imagePlaneProto := renderResponse.GetImagePlane()
	imagePlane := image.NewPlaneFromProto(imagePlaneProto)
	deserializeSpan.Finish()

	w.Header().Set("Content-Type", "image/png")

	if err := png.Encode(w, imagePlane.ToNRGBA()); err != nil {
		return errors.Wrap(err, "Failed to encode png")
	}

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/scenes/{scene_id}/nodes/{node_id}/render").Subrouter()

	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
}
