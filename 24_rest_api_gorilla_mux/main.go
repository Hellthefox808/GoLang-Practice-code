// Package main implements a production-grade, concurrency-safe RESTful API
// using Gorilla Mux, structured response envelopes, sync.RWMutex memory persistence,
// semantic HTTP status codes, and comprehensive validation guards.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// StandardAPIResponse enforces a unified JSON envelope across all endpoints.
type StandardAPIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

type Course struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	PriceUSD float64 `json:"price_usd"`
	Author   *Author `json:"author"`
}

// CourseRepository provides thread-safe in-memory CRUD operations.
type CourseRepository struct {
	mu      sync.RWMutex
	courses map[string]Course
	counter int64
}

func NewCourseRepository() *CourseRepository {
	repo := &CourseRepository{
		courses: make(map[string]Course),
		counter: 100,
	}
	// Seed initial items safely
	repo.courses["101"] = Course{
		ID:       "101",
		Title:    "Go Systems & Concurrency",
		PriceUSD: 99.00,
		Author:   &Author{FullName: "Ravi Ranjan Singh", Website: "https://github.com/Hellthefox808"},
	}
	repo.courses["102"] = Course{
		ID:       "102",
		Title:    "Cloud Native Microservices with Kubernetes",
		PriceUSD: 149.00,
		Author:   &Author{FullName: "Ravi Ranjan Singh", Website: "https://github.com/Hellthefox808"},
	}
	return repo
}

func (r *CourseRepository) GetAll() []Course {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]Course, 0, len(r.courses))
	for _, c := range r.courses {
		items = append(items, c)
	}
	return items
}

func (r *CourseRepository) GetByID(id string) (Course, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	course, exists := r.courses[id]
	return course, exists
}

func (r *CourseRepository) Create(c Course) (Course, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Title duplicate validation
	for _, existing := range r.courses {
		if existing.Title == c.Title {
			return Course{}, errors.New("a course with this title already exists")
		}
	}

	r.counter++
	c.ID = strconv.FormatInt(r.counter, 10)
	r.courses[c.ID] = c
	return c, nil
}

func (r *CourseRepository) Update(id string, updated Course) (Course, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.courses[id]; !exists {
		return Course{}, false
	}

	updated.ID = id
	r.courses[id] = updated
	return updated, true
}

func (r *CourseRepository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.courses[id]; !exists {
		return false
	}
	delete(r.courses, id)
	return true
}

// Controller holds dependencies for routing handlers.
type Controller struct {
	repo *CourseRepository
}

func NewController(repo *CourseRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) writeJSON(w http.ResponseWriter, status int, response StandardAPIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func (c *Controller) HandleHome(w http.ResponseWriter, r *http.Request) {
	c.writeJSON(w, http.StatusOK, StandardAPIResponse{
		Success: true,
		Message: "Welcome to the Go Production REST API authored by Ravi Ranjan Singh",
	})
}

func (c *Controller) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	courses := c.repo.GetAll()
	c.writeJSON(w, http.StatusOK, StandardAPIResponse{
		Success: true,
		Data:    courses,
	})
}

func (c *Controller) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	course, found := c.repo.GetByID(id)
	if !found {
		c.writeJSON(w, http.StatusNotFound, StandardAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Course with ID %q was not found", id),
		})
		return
	}

	c.writeJSON(w, http.StatusOK, StandardAPIResponse{
		Success: true,
		Data:    course,
	})
}

func (c *Controller) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		c.writeJSON(w, http.StatusBadRequest, StandardAPIResponse{
			Success: false,
			Error:   "Request body cannot be nil",
		})
		return
	}

	var newCourse Course
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newCourse); err != nil {
		c.writeJSON(w, http.StatusBadRequest, StandardAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Malformed JSON payload: %v", err),
		})
		return
	}

	if newCourse.Title == "" {
		c.writeJSON(w, http.StatusUnprocessableEntity, StandardAPIResponse{
			Success: false,
			Error:   "Field 'title' is mandatory",
		})
		return
	}

	created, err := c.repo.Create(newCourse)
	if err != nil {
		c.writeJSON(w, http.StatusConflict, StandardAPIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.writeJSON(w, http.StatusCreated, StandardAPIResponse{
		Success: true,
		Message: "Course registered successfully",
		Data:    created,
	})
}

func (c *Controller) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var payload Course
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		c.writeJSON(w, http.StatusBadRequest, StandardAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Invalid update payload: %v", err),
		})
		return
	}

	updated, ok := c.repo.Update(id, payload)
	if !ok {
		c.writeJSON(w, http.StatusNotFound, StandardAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Course ID %q not found for update", id),
		})
		return
	}

	c.writeJSON(w, http.StatusOK, StandardAPIResponse{
		Success: true,
		Message: "Course updated successfully",
		Data:    updated,
	})
}

func (c *Controller) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleted := c.repo.Delete(id)
	if !deleted {
		c.writeJSON(w, http.StatusNotFound, StandardAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Course ID %q not found for deletion", id),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// BuildServerRouter wires up endpoints and middleware.
func BuildServerRouter(ctrl *Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", ctrl.HandleHome).Methods(http.MethodGet)
	r.HandleFunc("/api/courses", ctrl.HandleGetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/courses/{id}", ctrl.HandleGetByID).Methods(http.MethodGet)
	r.HandleFunc("/api/courses", ctrl.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/api/courses/{id}", ctrl.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/api/courses/{id}", ctrl.HandleDelete).Methods(http.MethodDelete)

	return r
}

func main() {
	fmt.Println("=== 24: Enterprise REST API with Gorilla Mux ===")

	repo := NewCourseRepository()
	ctrl := NewController(repo)
	router := BuildServerRouter(ctrl)

	port := "4000"
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("REST API server listening at http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
