package app

import (
	_ "UchetUsers/docs"
	"UchetUsers/internal/configs"
	"UchetUsers/internal/handlers"
	"UchetUsers/internal/repository"
	"UchetUsers/internal/services"
	"UchetUsers/pkg/logger"
	"UchetUsers/pkg/postgres"
	"UchetUsers/pkg/server"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func Run() {
	v := validator.New()

	log, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	if err = InitConfig(); err != nil {
		log.Fatalf("Ошибка иницилизации конфига: %v", err)
	}

	if err = godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	db, err := postgres.ClientPostgres(log, configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.name"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	userRepo := repository.NewPostgresRepository(db, log)
	userService := services.NewUserService(userRepo, log, v)
	userHandler := handlers.NewUserHandler(userService, log)

	serv := new(server.Server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err = serv.RunServer(viper.GetString("serv.port"), userHandler.InitRoutes(log)); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	log.Info("Сервер запущен!")

	<-quit
	log.Info("Выключение сервера")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = serv.Shutdown(ctx); err != nil {
		log.Errorf("Ошибка при завершении сервера")
	}

	db.Close()
	log.Info("Сервер успешно выключен!")
}
