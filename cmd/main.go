package main

import (
	"TaskManager"
	"TaskManager/configs"
	"TaskManager/internal/handlers"
	"TaskManager/internal/repository"
	"TaskManager/internal/service"
	"TaskManager/pkg/database"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("init config err: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connection(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("database connection err: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	handler := handlers.NewTaskHandler(taskService)

	serv := new(TaskManager.Server)
	if err := serv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
