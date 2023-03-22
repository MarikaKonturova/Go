/*
*здесь мы реализуем логику подключения базы, а также будем хранить имена таблиц в константах
 */

package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
*для подключения к ДБ нам нужно 6 параметров, потому создадим тип
 */
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string //Имя базы данных
	SSLMode  string //режим SSL-подключения
}

/*
*создаём функцию, которая возвращает указатель на ДБ от sqlx и ошибку
 */
func NewPostgressDB(cfg Config) (*sqlx.DB, error) {
	/* @param=название движка
	   @param=строка с параметрами подключения
	*/
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	/*
	*проверяем подключение с помощью ф-ии Ping
	 */
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
