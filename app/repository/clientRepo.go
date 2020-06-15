package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbMap = Init()

func Init() *sql.DB {
	db, err := sql.Open("mysql",
		"root:my-secret-pw@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetUsers() (string, string) {
	var (
		id    string
		title string
	)
	db, err := sql.Open("mysql",
		"root:my-secret-pw@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	rows, qErr := db.Query("select id, title from usr")
	if qErr != nil {
		log.Fatal(qErr)
	}

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title)
	}

	defer db.Close()
	defer rows.Close()
	return id, title
}
