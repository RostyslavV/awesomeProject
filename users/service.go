package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Service - creates new structure.
type Service struct {
	users DB
}

// New - creates new means of service.
func New(users DB) *Service {
	return &Service{users: users}
}

// Create - creates user into DB.
func (s *Service) Create(ctx context.Context, name, email, password string) error {
	return s.users.Create(ctx, User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  []byte(password),
		CreatedAt: time.Now(),
	})
}

// Delete - deletes user from DB.
func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.users.Delete(ctx, id)
}

// Get - returns user from DB.
func (s *Service) Get(ctx context.Context, id uuid.UUID) (User, error) {
	return s.users.Get(ctx, id)
}

// Update - updates user in DB.
func (s *Service) Update(ctx context.Context, id uuid.UUID, name, email, password string) error {
	return s.users.Update(ctx, User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  []byte(password),
		CreatedAt: time.Now(),
	})
}
