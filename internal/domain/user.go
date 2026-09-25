package domain

import "time"

// User represents the user entity in our clean architecture.
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserRepository defines the input/output port for user data access.
type UserRepository interface {
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
}

// UserService defines the input/output port for user business logic (use cases).
type UserService interface {
	Register(email, password string) (*User, error)
	Login(email, password string) (string, error)
	GetUser(id string) (*User, error)
}
