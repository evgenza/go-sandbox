package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("Таблица успешно создана или уже существует.")

	// Вставка данных
	insertUserQuery := `
	INSERT INTO users (name, age) VALUES (?, ?);
	`
	_, err = db.Exec(insertUserQuery, "Alice", 25)
	if err != nil {
		log.Fatalf("Ошибка вставки данных: %v", err)
	}
	_, err = db.Exec(insertUserQuery, "Bob", 30)
	if err != nil {
		log.Fatalf("Ошибка вставки данных: %v", err)
	}
	fmt.Println("Данные успешно вставлены.")

	// Чтение данных
	readUsersQuery := `
	SELECT id, name, age FROM users;
	`
	rows, err := db.Query(readUsersQuery)
	if err != nil {
		log.Fatalf("Ошибка чтения данных: %v", err)
	}
	defer rows.Close()

	fmt.Println("Список пользователей:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("Ошибка сканирования строки: %v", err)
		}
		fmt.Printf("ID: %d, Имя: %s, Возраст: %d\n", id, name, age)
	}

	// Обновление данных
	updateUserQuery := `
	UPDATE users SET age = ? WHERE name = ?;
	`
	_, err = db.Exec(updateUserQuery, 28, "Alice")
	if err != nil {
		log.Fatalf("Ошибка обновления данных: %v", err)
	}
	fmt.Println("Данные обновлены.")

	// Удаление данных
	deleteUserQuery := `
	DELETE FROM users WHERE name = ?;
	`
	_, err = db.Exec(deleteUserQuery, "Bob")
	if err != nil {
		log.Fatalf("Ошибка удаления данных: %v", err)
	}
	fmt.Println("Данные удалены.")
}
