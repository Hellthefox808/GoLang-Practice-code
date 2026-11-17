// Package controller handles HTTP routing, request parsing, validation,
// and delegates to database operations with resilient error handling.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieController struct {
	collection *mongo.Collection
}

func NewMovieController(collection *mongo.Collection) *MovieController {
	return &MovieController{collection: collection}
}

func (c *MovieController) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (c *MovieController) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if c.collection == nil {
		c.writeJSON(w, http.StatusOK, []model.Movie{})
		return
	}

	cur, err := c.collection.Find(ctx, bson.D{{}})
	if err != nil {
		c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		return
	}
	defer cur.Close(ctx)

	var movies []model.Movie
	for cur.Next(ctx) {
		var m model.Movie
		if err := cur.Decode(&m); err != nil {
			c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Decode failure"})
			return
		}
		movies = append(movies, m)
	}

	if err := cur.Err(); err != nil {
		c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Cursor iteration error"})
		return
	}

	if movies == nil {
		movies = []model.Movie{}
	}

	c.writeJSON(w, http.StatusOK, movies)
}

func (c *MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		c.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if movie.Title == "" {
		c.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "Movie title is required"})
		return
	}

	movie.ID = primitive.NewObjectID()

	if c.collection != nil {
		_, err := c.collection.InsertOne(ctx, movie)
		if err != nil {
			c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to insert movie"})
			return
		}
	}

	c.writeJSON(w, http.StatusCreated, movie)
}

func (c *MovieController) MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	idStr := vars["id"]

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ObjectID format"})
		return
	}

	if c.collection != nil {
		filter := bson.M{"_id": objID}
		update := bson.M{"$set": bson.M{"watched": true}}
		res, err := c.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Update failed"})
			return
		}
		if res.MatchedCount == 0 {
			c.writeJSON(w, http.StatusNotFound, map[string]string{"error": "Movie not found"})
			return
		}
	}

	c.writeJSON(w, http.StatusOK, map[string]string{"status": "marked as watched", "id": idStr})
}

func (c *MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	idStr := vars["id"]

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ObjectID format"})
		return
	}

	if c.collection != nil {
		filter := bson.M{"_id": objID}
		res, err := c.collection.DeleteOne(ctx, filter)
		if err != nil {
			c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Delete failed"})
			return
		}
		if res.DeletedCount == 0 {
			c.writeJSON(w, http.StatusNotFound, map[string]string{"error": "Movie not found"})
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *MovieController) DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var count int64
	if c.collection != nil {
		res, err := c.collection.DeleteMany(ctx, bson.D{{}}, options.Delete())
		if err != nil {
			c.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Bulk delete failed"})
			return
		}
		count = res.DeletedCount
	}

	c.writeJSON(w, http.StatusOK, map[string]any{"deleted_count": count})
}
