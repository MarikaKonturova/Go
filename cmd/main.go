package main

import (
	"log"

	"github.com/MarikaKonturova/todo-go-app"
	handler "github.com/MarikaKonturova/todo-go-app/pkg/handlers"
	"github.com/MarikaKonturova/todo-go-app/pkg/repository"
	"github.com/MarikaKonturova/todo-go-app/pkg/service"
)

/* точка запуска программы */
func main() {
	/*
	 *инициализируем экземпляр сервера с помощью ключевого слова new
	 */
	srv := new(todo.Server)

	/*
	*объявляем все зависимости в нужном порядке
	 */
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
