package main

import (
	"first/connection"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := connection.Connect(os.Getenv("DBNAME"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("SSLMODE"))
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	us := User{Name: "Jane Doe", Email: "jhon@gmail.com"}
	_, err = db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", us.Name, us.Email)
	if err != nil {
		log.Fatalln(err)
	}
	// get all users from database
	var user User
	rows, _ := db.Queryx("SELECT * FROM users ")
	for rows.Next() {
		rows.StructScan(&user)
		log.Println(user)
	}

}
