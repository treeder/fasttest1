package models

import "github.com/treeder/firetils"

type Vote struct {
	firetils.Timestamped
	firetils.Firestored
	firetils.IDed
	firetils.Owned

	PostID string `json:"post_id" firestore:"post_id"`
	Count  int    `json:"count" firestore:"count"`
}
