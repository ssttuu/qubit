package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"github.com/stupschwartz/qubit/core/image"
	pb "github.com/stupschwartz/qubit/compute/protos/compute"
	"net/http"

	"golang.org/x/net/context"
	"image/png"
	"strconv"
	"github.com/pkg/errors"
	"cloud.google.com/go/trace"
	"io"
	"fmt"
)

func PostHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PostHandler")
	fmt.Println("PostHandler")
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

	fmt.Println("Get Stream")
	stream, err := e.ComputeClient.Render(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get stream")
	}

	fmt.Println("Send Stream")

	var i int64
	for i = 0; i < height; i++ {
		fmt.Printf("%v;", i)
		renderRequest := &pb.RenderRequest{SceneId: sceneId, NodeId: nodeId, BoundingBox: &pb.BoundingBox{StartX: 0, StartY: i, EndX: width, EndY: i + 1 }}
		err := stream.Send(renderRequest)
		if err != nil {
			return errors.Wrap(err, "Failed to send stream")
		}
	}
	stream.CloseSend()

	imagePlane := image.NewRGBZeroPlane(width, height)

	deserializeSpan := span.NewChild("Get gRPC Response")
	fmt.Println(deserializeSpan)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "Failed to receive response")
		}

		imagePlanePart := image.NewPlaneFromProto(in.GetImagePlane())
		partBBox := in.GetBoundingBox()
		imagePlane.Merge(&imagePlanePart, partBBox.GetStartX(), partBBox.GetStartY())
	}

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
