package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"g
	dbname   = "BrotHER"
	password = "1234"
	host     = "localhost"
	port     = 5433
)

func main() {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", user, dbname, password, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при открытии соединения:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
	printAllBrothers(db)
	printMe(db)
}
func printAllBrothers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM brothers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var family string
		var name string
		var surname string
		if err := rows.Scan(&id, &family, &name, &surname); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, family, name, surname)
	}
}
func printMe(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM brothers WHERE family = 'Borisov';")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var family string
		var name string
		var surname string
		if err := rows.Scan(&id, &family, &name, &surname); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, family, name, surname)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
