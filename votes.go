package main

import (
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi"
	"github.com/treeder/fasttest1/models"
	"github.com/treeder/firetils"
	"github.com/treeder/gotils"
)

func postVote(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	postID := chi.URLParam(r, "id")
	p := &models.Vote{}
	err := gotils.ParseJSONReader(r.Body, p)
	if err != nil {
		return errors.New("Invalid json")
	}

	// // validate
	// if p.Title == "" {
	// 	return errors.New("Post requires a title")
	// }

	// only one vote object per post per user
	var vote *models.Vote
	v, err := firetils.GetOneByQuery2(ctx, firetils.Collection(fs, "votes").Where("user_id", "==", firetils.UserID(ctx)).Where("post_id", "==", postID), p)
	if err != nil {
		if !errors.Is(err, gotils.ErrNotFound) {
			return err
		}
		vote = p
		vote.PostID = postID
	} else {
		vote = v.(*models.Vote)
	}
	vote.Count += p.Count
	v, err = firetils.Save(ctx, fs, "votes", vote)
	if err != nil {
		return err
	}
	p = v.(*models.Vote)

	// TOOD: should cache this on a per post basis for a second or two since can only update once per second:
	// https://firebase.google.com/docs/firestore/manage-data/add-data#increment_a_numeric_value

	postRef := fs.Collection("posts").Doc(postID)
	_, err = postRef.Update(ctx, []firestore.Update{
		{Path: "votes", Value: firestore.Increment(p.Count)},
	})
	if err != nil {
		return err
	}

	gotils.WriteObject(w, http.StatusOK, p)
	return nil
}

// func getPost(w http.ResponseWriter, r *http.Request) error {
// 	ctx := r.Context()
// 	id := chi.URLParam(r, "id")
// 	p := &models.Post{}
// 	err := firetils.GetByID(ctx, fs, "posts", id, p)
// 	if err != nil {
// 		return err
// 	}
// 	gotils.WriteObject(w, http.StatusOK, p)
// 	return nil
// }

// func getPosts(w http.ResponseWriter, r *http.Request) error {
// 	ctx := r.Context()
// 	p := &models.Post{}
// 	ps, err := firetils.GetAllByQuery2(ctx, firetils.Collection(fs, "posts").Query, p)
// 	if err != nil {
// 		return err
// 	}
// 	gotils.WriteObject(w, http.StatusOK, map[string]interface{}{"posts": ps})
// 	return nil
// }
// func deletePost(w http.ResponseWriter, r *http.Request) error {
// 	ctx := r.Context()
// 	id := chi.URLParam(r, "id")
// 	err := firetils.Delete(ctx, fs, "posts", id)
// 	if err != nil {
// 		return err
// 	}
// 	gotils.WriteMessage(w, http.StatusOK, "deleted")
// 	return nil
// }
