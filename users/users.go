package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User - contains fields that describe user entity.
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

// DB - provides access to users data base.
type DB interface {
	Create(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (User, error)
	Update(ctx context.Context, user User) error
}
