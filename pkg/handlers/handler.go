package handler

import (
	"github.com/MarikaKonturova/todo-go-app/pkg/service"
	"github.com/gin-gonic/gin"
)

/*
*директория для имплементации handler-ов
 */

// пока пустая, но в дальнейшем добавим поля
type Handler struct {
	/*
	*наши обработчики будут вызывать методы сервисов, потому добавим в нашу структуру указатель на сервис
	 */
	services *service.Service
}

/*
*создадим конструктор, в который будем внедрять зависимости сервиса
 */
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

/*
*Объявляем метод InitRoutes, который будет инициализировать все наши endpoint-ы

* Для разработки Api будем использовать популярный фреймворк gin, он популярен своим синтаксисом, а также производительностью, разработчики заявляют, что он в 40 раз быстрее стандартного .netHttp

 */

func (h *Handler) InitRoutes() *gin.Engine {
	//инициализация роутера
	router := gin.New()
	//объявляем методы, сгруппировав их по маршруту
	//3 и добавим пустые обработчики роутов
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-In", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			// динамический параметр - фишка библиотеки
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}

	}
	return router
}
