/*
*здесь объявим заготовки интерфейсов для наших сущностей
 */
package repository

import "github.com/jmoiron/sqlx"

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
*поскольку наши репозитории должны работать с ДББ передаём аргументы в конструктор
*/

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
