package main

import (
				"fmt"
				"log"
				"net/http"
				"github.com/gorilla/mux"
				"database/sql"
				_ "github.com/lib/pq"
				"os"
)

var db *sql.DB

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)
func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}



func main() {
	initDb()
  defer db.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8081", router))
	// fmt.Println("Hello World!")
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	config[dbhost], config[dbport],
	config[dbuser], config[dbpass], config[dbname])
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
			panic(err)
	}
	err = db.Ping()
	if err != nil {
			panic(err)
	}
	fmt.Println("Successfully connected!")
}



func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
			panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
			panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
			panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
			panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
			panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}