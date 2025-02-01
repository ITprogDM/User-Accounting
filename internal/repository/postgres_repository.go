package repository

import (
	"UchetUsers/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgresRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewPostgresRepository(db *pgxpool.Pool, logger *logrus.Logger) *PostgresRepository {
	return &PostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (r *PostgresRepository) CreateUser(ctx context.Context, u models.User) error {
	query := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, query, u.Name, u.Email, u.Age)
	if err != nil {
		r.logger.Info("Неверный запрос для создания пользователя к БД")
		return err
	}

	r.logger.Infof("Пользователь %s успешно создан", u.Name)
	return nil
}

func (r *PostgresRepository) GetUser(ctx context.Context, id int) (*models.User, error) {
	var u models.User

	query := "SELECT id, name, email, age FROM users WHERE id = $1"

	err := r.db.QueryRow(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	if err != nil {
		r.logger.Info("Неверный запрос для поиска к БД")
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, u models.User) error {
	query := "UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4"

	_, err := r.db.Exec(ctx, query, u.Name, u.Email, u.Age, u.ID)
	if err != nil {
		r.logger.Info("Неверный запрос для обновления пользователя к БД")
		return err
	}

	r.logger.Infof("Данные пользователя %s обновлены", u.Name)
	return err
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Info("Неверный запрос для удаления пользователя из БД")
		return err
	}

	r.logger.Infof("Пользователь c номером id: %d - успешно удалён!", id)
	return nil
}

func (r *PostgresRepository) IsUniqueEmail(ctx context.Context, email string) (bool, error) {
	var count int

	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	err := r.db.QueryRow(ctx, query, email).Scan(&count)
	if err != nil {
		r.logger.WithError(err).Error("Ошибка при выполнении запроса  для проверки уникальности email")
		return false, err
	}

	return count == 0, nil
}
