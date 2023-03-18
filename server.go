package todo

import (
	"context"
	"net/http"
	"time"
)

/*
*Здесь мы объявляем структуру данных, которую будем использовать для запуска http-сервера */
type Server struct {
	httpServer *http.Server
}

/*
*Эта структура является небольшой абстракцией над сервером из пакета http и имеет всего одно поле - указатель на эту структуру;

*У сервера будет 2 метода: запуск и остановка работы

*в методе запуска мы инкапсулируем значения для таких значений как MaxHeaderBytes, read&write Timeout, установив значение на 1 мб и 10 секунд соотвтетственно также мы будем принимать значения для порта, на котором будет запускаться наш сервер

*возвращать мы будем стандартный метод httpServer-а listenAndServe(), который под капотом запускает бесконечный цикл for и слушает все входящие запросы для последующей обработки
 */

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
        Handler: handler,
		MaxHeaderBytes: 1 << 20, //1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

/*
*Также объявим метод shutdown, который будем использовать при выходе из приложения
 */
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
