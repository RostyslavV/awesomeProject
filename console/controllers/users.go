package controllers

import (
	"net/http"

	"github.com/google/uuid"

	"awesomeProject/users"
	"encoding/json"
)

// Users creates new structure.
type Users struct {
	users *users.Service
}

// NewUsers creates new means of users.
func NewUsers(users *users.Service) *Users {
	return &Users{users: users}
}

// CreateRequest - creates Create structure
type CreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create - creates user into server
func (u *Users)Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := u.users.Create(ctx, req.Name, req.Email, req.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

// GetRequest creates Get structure
type GetRequest struct {
	ID string `json:"id"`
}

// Get - returns user from server
func (u *Users) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req GetRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := u.users.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// DeleteRequest - creates Delete structure
type DeleteRequest struct {
	ID string `json:"id"`
}

// Delete - deletes user from server
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req DeleteRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := u.users.Delete(ctx, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// UpdateRequest - creates Update structure
type UpdateRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Update - updates user in server
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req UpdateRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := u.users.Update(ctx, id, req.Name, req.Email, req.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
