// Creates an HTTP server at :1357 to receive the requests from the Angular application at /assets
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Passwd     string    `json:"passwd"`
}

func main() {
	SetupDB()

	fs := http.FileServer(http.Dir("assets"))

	router := mux.NewRouter()
	
	router.HandleFunc("/user", UserCreate).Methods("POST")
	router.PathPrefix("/").Handler(fs)

	log.Println("Server listing at :1357")
	log.Fatal(http.ListenAndServe(":1357", router))
}

// Creates a new SQLite file 'gowels.db' and creates the 'reviews' table
func SetupDB() {
	var err error

	db, err = sql.Open("postgres", "user=postgres dbname=melancholia  password='lucas' sslmode=disable")
	if err != nil {
		log.Printf("---Open PG---")
		log.Fatal(err)
	}

	createTable := `CREATE TABLE users_angular(username VARCHAR(30), email VARCHAR(50), passwd VARCHAR(50))`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Printf("%q here: %s\n", err, createTable)
		return
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(201) 
	
	log.Printf("Usercreating")

	decoder := json.NewDecoder(r.Body)
	var user User

	err := decoder.Decode(&user)
	if err != nil {
		log.Printf("decoder.Decode(&user) -error")
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Begin() -error")
		log.Fatal(err)
	}
	
	fmt.Println(user.Username, user.Passwd)

	stmt, err := tx.Prepare("INSERT into users_angular(username, email, passwd) values($1, $2, $3);")
	if err != nil {
		log.Printf("Prepare() -error")
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Email, user.Passwd)
	if err != nil {
		log.Printf("Exec() -error")
		log.Fatal(err)
	}
	tx.Commit()
}