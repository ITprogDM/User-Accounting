package services

import (
	"UchetUsers/internal/models"
	"UchetUsers/internal/repository"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	rep       repository.UserRepository
	logger    *logrus.Logger
	validator *validator.Validate
}

func NewUserService(rep repository.UserRepository, logger *logrus.Logger, validator *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		rep:       rep,
		logger:    logger,
		validator: validator,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, u models.User) error {
	if err := s.validator.Struct(u); err != nil {
		s.logger.WithError(err).Error("Ошибка валидации данных пользователя")
		return fmt.Errorf("Ошибка валидации: %w", err)
	}

	IsUnique, err := s.rep.IsUniqueEmail(ctx, u.Email)
	if err != nil {
		s.logger.WithError(err).Error("Ошибка базы данных")
		return fmt.Errorf("Ошибка базы данных при проверке уникальности email: %w", err)
	}
	if !IsUnique {
		s.logger.Warnf("Пользователь с email %s уже существует!", u.Email)
		return fmt.Errorf("Пользовтель с email %s уже зарегистрирован", u.Email)
	}

	if err = s.rep.CreateUser(ctx, u); err != nil {
		s.logger.WithError(err).Error("Ошибка создания пользователя в базе данных")
		return fmt.Errorf("Ошибка создания пользователя: %w", err)
	}

	s.logger.Infof("Пользователь с email %s успешно создан!", u.Email)
	return nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, err := s.rep.GetUser(ctx, id)
	if err != nil {
		s.logger.WithError(err).Error("Ошибка поиска пользователя в базе данных")
		return nil, fmt.Errorf("Данный пользователь с id: %s не существует", id)
	}

	return user, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, id int, name, email string, age int) error {
	user := &models.User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}

	err := s.validator.Struct(user)
	if err != nil {
		s.logger.WithError(err).Error("Ошибка валидации данных пользователя")
		return fmt.Errorf("Ошибка валидации: %w", err)
	}

	isUnique, err := s.rep.IsUniqueEmail(ctx, email)
	if err != nil {
		s.logger.WithError(err).Error("Ошибка базы данных")
		return fmt.Errorf("Ошибка базы данных при проверке уникальности email: %w", err)
	}
	if !isUnique {
		s.logger.Warnf("Пользователь с таким email %s уже существует!", email)
		return fmt.Errorf("Email %s уже занят", email)
	}

	if err = s.rep.UpdateUser(ctx, *user); err != nil {
		s.logger.WithError(err).Error("Ошибка обновления данных пользователя")
		return fmt.Errorf("Ошибка обновления данных: %w", err)
	}

	s.logger.Infof("Данные пользователя с id: %d успешно обновлены", id)
	return nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	if err := s.rep.DeleteUser(ctx, id); err != nil {
		s.logger.WithError(err).Error("Ошибка удаления пользователя")
		return fmt.Errorf("Ошибка, пользователь не удалён: %w", err)
	}

	s.logger.Infof("Пользовтаель с id %d успешно удалён!", id)
	return nil
}
