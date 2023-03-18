package todo

/*
* Файл для структур пользователя и сущностей Todo (User)
 */
/*
* Поля стурктуры User-а полностью идентичны с полями БД, только поле password в БД называется passwordHash

*также указываем на наших полях json-теги, чтобы корректно принимать и выводить данные при http-запросах
 */

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username int    `json:"username"`
	Password int    `json:"password"`
}
