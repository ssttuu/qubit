package scenes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
	"github.com/satori/go.uuid"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"cloud.google.com/go/trace"
	"github.com/stupschwartz/qubit/core/scene"
)


func PutScene(ctx context.Context, sceneKey *datastore.Key, s *scene.Scene, e *env.Env) error {
	span := trace.FromContext(ctx).NewChild("PutScene")
	defer span.Finish()

	if _, err := e.DatastoreClient.Put(ctx, sceneKey, s); err != nil {
		return errors.Wrapf(err, "Failed to put node with digest %v", s.Id)
	}

	return nil
}


func GetAllHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("GetAllHandler")
	defer span.Finish()

	var scenes []*scene.Scene
	_, err := env.DatastoreClient.GetAll(ctx, datastore.NewQuery("Scene"), &scenes)
	if err != nil {
		return errors.Wrap(err, "Could not get all")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(&scenes)
	fmt.Fprint(w, string(jsonData))

	return nil
}

func GetHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("GetHandler")
	defer span.Finish()

	vars := mux.Vars(r)

	sceneId := vars["scene_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	var existingScene scene.Scene
	if err := env.DatastoreClient.Get(ctx, sceneKey, &existingScene); err != nil {
		return errors.Wrap(err, "Could not get datastore entity")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	sceneAsJson, err := json.Marshal(&existingScene)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json")
	}

	fmt.Fprintf(w, string(sceneAsJson))

	return nil
}

type RequestBody struct {
	Scene   *scene.Scene `json:"scene"`
	Params params.Parameters `json:"params"`
}

func PostHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PostHandler")
	defer span.Finish()

	sceneUuid := uuid.NewV4()
	sceneKey := datastore.NameKey("Scene", sceneUuid.String(), nil)

	requestBody := RequestBody{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		return errors.Wrap(err, "Failed to decode json request body")
	}

	newScene := requestBody.Scene

	defer r.Body.Close()

	newScene.Id = sceneUuid.String()
	newScene.Version = 0
	PutScene(ctx, sceneKey, newScene, env)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	sceneAsJson, err := json.Marshal(newScene)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal json")
	}

	fmt.Fprintf(w, string(sceneAsJson))

	return nil
}

func PutHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PutHandler")
	defer span.Finish()

	vars := mux.Vars(r)
	sceneId := vars["scene_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	decoder := json.NewDecoder(r.Body)

	requestBody := RequestBody{}

	if err := decoder.Decode(&requestBody); err != nil {
		return errors.Wrap(err, "Failed to decode request body")
	}

	newScene := requestBody.Scene

	_, err := env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingScene scene.Scene
		if err := tx.Get(sceneKey, &existingScene); err != nil {
			return errors.Wrap(err, "Failed to get scene in tx")
		}
		existingScene.Version += 1
		existingScene.Name = newScene.Name
		PutScene(ctx, sceneKey, &existingScene, env)

		_, err := tx.Put(sceneKey, &existingScene)
		if err != nil {
			return errors.Wrap(err, "Failed to put scene in tx")
		}
		return nil
	})

	if err != nil {
		return errors.Wrap(err, "Failed to update scene")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	sceneAsJson, err := json.Marshal(&newScene)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal json")
	}

	fmt.Fprintf(w, string(sceneAsJson))

	return nil
}

func DeleteAllHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("DeleteAllHandler")
	defer span.Finish()

	var scenes interface{}
	sceneIds, err := e.DatastoreClient.GetAll(ctx, datastore.NewQuery("Scene").KeysOnly(), scenes)
	if err != nil {
		return errors.Wrap(err, "Failed to get all keys only")
	}

	err = e.DatastoreClient.DeleteMulti(ctx, sceneIds)
	if err != nil {
		return errors.Wrap(err, "Failed to delete multi")
	}

	return nil
}

func DeleteHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("DeleteHandler")
	defer span.Finish()

	return nil
}


func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/scenes/").Subrouter()

	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
	s.Handle("/{scene_id}/", handler.Handler{environ, GetHandler}).Methods("GET")

	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
	s.Handle("/{scene_id}/", handler.Handler{environ, PutHandler}).Methods("PUT")
	s.Handle("/DELETE/", handler.Handler{environ, DeleteAllHandler}).Methods("DELETE")
	s.Handle("/{scene_id}/", handler.Handler{environ, DeleteHandler}).Methods("DELETE")
}
