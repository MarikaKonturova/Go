/* 
*используется при откате текующей версии на предыдущую, то есть если мы в файлу .up добавляем новые таблицы или изменяем существующие, то в .down мы должны откатывать все изменения из файла .up; в нашем случае нужно просто дропнуть все таблицы
 */
DROP TABLE lists_items;

DROP TABLE users_lists;

DROP TABLE todo_lists;

DROP TABLE users;

DROP TABLE todo_items;