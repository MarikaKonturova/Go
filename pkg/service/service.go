/*
*здесь объявим заготовки интерфейсов для наших сущностей
 */
package service

import "github.com/MarikaKonturova/todo-go-app/pkg/repository"

/*
*называем интерфес согласно их доменной зоне, то есть участка бизнес-логики приложения, за которую они отвечают
 */
type Authorization interface{}
type TodoList interface{}
type TodoItem interface{}

/* 
*собираем структуру сервиса, которая собирает все наши интерфейсы вместе
*/

type Service struct {
	Authorization
	TodoList
	TodoItem
}
/* 
*объявим конструктор

*сервис будет обращаться к БД => реализуем внедрение зависимостей через передачу параметра repos с указателем на структуру репозитория из пакета репозиторий
*/

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
