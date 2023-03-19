/*
*здесь объявим заготовки интерфейсов для наших сущностей
 */
package repository

/*
 *называем интерфес согласно их доменной зоне, то есть участка бизнес-логики приложения, за которую они отвечают
 */
type Authorization interface{}
type TodoList interface{}
type TodoItem interface{}

/*
 *собираем структуру репозитория, которая собирает все наши интерфейсы вместе
 */

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

/*
 *объявим конструктор
 */

func NewRepository() *Repository {
	return &Repository{}
}
