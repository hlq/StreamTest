package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "src/github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	//sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//{ // Create a new table
	//	query := `
	//        CREATE TABLE myusers (
	//            id INT AUTO_INCREMENT,
	//            username TEXT NOT NULL,
	//            password TEXT NOT NULL,
	//            created_at DATETIME,
	//            PRIMARY KEY (id)
	//        );`
	//
	//	if _, err := db.Exec(query); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO myusers (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM myusers WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}
	
		rows, err := db.Query(`SELECT id, username, password, created_at FROM myusers`)
		if err != nil{
			log.Fatal(err)
		}
		
		defer rows.Close()
		
		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("%#v", users)
		
	}

	{
		_, err := db.Exec(`DELETE FROM myusers WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}

}
