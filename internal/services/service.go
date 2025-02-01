package services

import (
	"UchetUsers/internal/models"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, u models.User) error
	GetUser(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, id int, name, email string, age int) error
	DeleteUser(ctx context.Context, id int) error
}
