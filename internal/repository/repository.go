package repository

import (
	"UchetUsers/internal/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u models.User) error
	GetUser(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, u models.User) error
	DeleteUser(ctx context.Context, id int) error
	IsUniqueEmail(ctx context.Context, email string) (bool, error)
}
