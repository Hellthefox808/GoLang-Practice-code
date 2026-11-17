// Package model defines domain entities and BSON/JSON mapping tags.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie represents a movie title in the watchlist.
type Movie struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Watched   bool               `json:"watched" bson:"watched"`
	Rating    float64            `json:"rating,omitempty" bson:"rating,omitempty"`
}
