package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/treeder/fasttest1/models"
	"github.com/treeder/firetils"
	"github.com/treeder/gotils"
)

func postPost(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	p := &models.Post{}
	err := gotils.ParseJSONReader(r.Body, p)
	if err != nil {
		return errors.New("Invalid json")
	}

	// validate
	if p.Title == "" {
		return errors.New("Post requires a title")
	}

	v, err := firetils.Save(ctx, fs, "posts", p)
	if err != nil {
		return err
	}
	p = v.(*models.Post)
	gotils.WriteObject(w, http.StatusOK, p)
	return nil
}

func getPost(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	p := &models.Post{}
	err := firetils.GetByID(ctx, fs, "posts", id, p)
	if err != nil {
		return err
	}
	gotils.WriteObject(w, http.StatusOK, p)
	return nil
}

func getPosts(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	p := &models.Post{}
	ps, err := firetils.GetAllByQuery2(ctx, firetils.Collection(fs, "posts").Query, p)
	if err != nil {
		return err
	}
	gotils.WriteObject(w, http.StatusOK, map[string]interface{}{"posts": ps})
	return nil
}
func deletePost(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	err := firetils.Delete(ctx, fs, "posts", id)
	if err != nil {
		return err
	}
	gotils.WriteMessage(w, http.StatusOK, "deleted")
	return nil
}
