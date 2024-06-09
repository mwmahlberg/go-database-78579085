package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	host     string
	port     int
	dbname   string
	username string
	password string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host")
	flag.IntVar(&port, "port", 7899, "port")
	flag.StringVar(&dbname, "dbname", "NEWDB1", "dbname")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
}

func main() {
	flag.Parse()
	log.Println("started")
	connectionData := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, dbname)
	conn, err := sqlx.Open("postgres", connectionData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database")
	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("ping")
	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("close")
}
