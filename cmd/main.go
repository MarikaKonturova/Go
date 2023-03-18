package main

import (
	"log"

	"github.com/MarikaKonturova/todo-go-app"
	handler "github.com/MarikaKonturova/todo-go-app/pkg/handlers"
)

/* точка запуска программы */
func main() {
	/*
	 *инициализируем экземпляр сервера с помощью ключевого слова new
	 */
	srv := new(todo.Server)

	/*
		//1
			* а также запустим сервер на порту 8000 через метод run

			* если во время работы метод возвращает ошибку, тогда мы сделаем вызов метода Fatalf из стандартного пакета log, который выведет ошибку на экран и выйдет из приложения

			//2
			* теперь создадим экземпляр объекта handler из нашего пакета pkg/hander, а в метод run передадим ф-ю initRoutes

			*т.к. init Routes возвращает объект типа указателя на объект gin.Engine, кот. реализует интерфейс handler-а с пакета http - мы без проблем можем использовать его в кач-ве аргумента для нашего сервера

		//3
		* запускаем программу через go run cmd/main.go. 

		*Однако мы видим панику. Потому что для end-point-a нужен как минимум 1 обработчик - теперь создадим пустые обработчики и присвомим их к маршрутам

		//3
		*после добавления обработчиков запускаем сервер через команду go run cmd/main.go, и видим, что сервер запускается с набором описанных ендпоинтов
	*/
	//2
	handlers := new(handler.Handler)

	//1
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
