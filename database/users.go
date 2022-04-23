package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"awesomeProject/users"
)

// userDB - creates new structure.
type userDB struct {
	pool *pgxpool.Pool
}

// Create - creates user into DB.
func (u *userDB) Create(ctx context.Context, user users.User) error {
	passwordHast, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.pool.Exec(ctx, `INSERT INTO users(id, email, name, password, created_at) 
						VALUES ($1,$2,$3,$4,$5)`, user.ID, user.Email, user.Name, passwordHast, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Delete - deletes user from DB.
func (u *userDB) Delete(ctx context.Context, uuid uuid.UUID) error {
	_, err := u.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, uuid)
	if err != nil {
		return err
	}

	return nil
}

// Get - returns user from DB.
func (u *userDB) Get(ctx context.Context, uuid uuid.UUID) (user users.User, err error) {
	row := u.pool.QueryRow(ctx, `SELECT id, email, name, password, created_at FROM users WHERE id = $1`, uuid)

	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

// Update - updates user in DB.
func (u *userDB) Update(ctx context.Context, user users.User) error {
	_, err := u.pool.Exec(ctx, `UPDATE users SET email = $1, name = $2, password = $3 WHERE id = $4`, user.Email, user.Name, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}
