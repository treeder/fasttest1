package models

import "github.com/treeder/firetils"

type Post struct {
	firetils.Timestamped
	firetils.Firestored
	firetils.IDed

	Title string `json:"title" firestore:"title"`
	Body  string `json:"body" firestore:"body"`
}
