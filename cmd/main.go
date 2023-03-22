package main

import (
	"log"

	"github.com/MarikaKonturova/todo-go-app"
	handler "github.com/MarikaKonturova/todo-go-app/pkg/handlers"
	"github.com/MarikaKonturova/todo-go-app/pkg/repository"
	"github.com/MarikaKonturova/todo-go-app/pkg/service"
	"github.com/spf13/viper"
)

/* точка запуска программы */
func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s ", err.Error())
	}
	srv := new(todo.Server)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

/*
*создаём функцию для работы с конфигурацией
 */

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
