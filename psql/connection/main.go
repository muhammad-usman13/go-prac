package connection

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbname string, dbhost string, dbport string, dbuser string, dbpass string, sslMode string) *sqlx.DB {
	dbStr := "user=" + dbuser + " dbname=" + dbname + " host=" + dbhost + " sslmode=" + sslMode + " password=" + dbpass
	fmt.Println(dbStr)
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {

		log.Fatalln("", err)
	}
	return db
}
