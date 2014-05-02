package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func httpPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		var person_map map[string]interface{}

		if err := json.Unmarshal(body, &person_map); err != nil {
			panic(err)
		}

		database := Open()
		defer database.Close()

		// f := person_map["id"].(float64)
		fullname := fmt.Sprintf("%s %s", person_map["first_name"].(string), person_map["last_name"].(string))

		p := Person{
			id:      int(person_map["id"].(float64)),
			name:    fullname,
			email:   person_map["email"].(string),
			updates: 0}

		database.Insert(p)

		fmt.Printf("Person updated: %s\n", string(fullname))
	}
}

func httpGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		database := Open()

		defer database.Close()
		database.ListPeople(w)
	}
}

const PORT = ":4567"

func main() {
	http.HandleFunc("/person", httpPostHandler)
	http.HandleFunc("/people", httpGetHandler)

	error := http.ListenAndServe(PORT, nil)
	if error != nil {
		panic(error)
	}
}

func init() {
	// Use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
}


// database specific code below here

type Person struct {
	id      int
	name    string
	email   string
	updates int
}

type Database struct {
	database sql.DB
}

func Open() Database {
	database_filename := "people.db"

	db, err := sql.Open("sqlite3", database_filename)
	if err != nil {
		fmt.Println("Unable to open the database: %s", err)
		os.Exit(1)
	}

	db.Exec(`CREATE TABLE people(id INTEGER PRIMARY KEY ON CONFLICT REPLACE, name VARCHAR(20),
        email TEXT, updates INTEGER DEFAULT 1);`)

	return Database{database: *db}
}

func (db Database) Insert(p Person) {
	query := "SELECT updates FROM people where id = ?"

	rows, err := db.database.Query(query, p.id)
	if err != nil {
		log.Fatal(err)
	}

	var updates int

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&updates)
	}

	updates++

	fmt.Sprintln("updates => %d", updates)

	insert := "INSERT INTO people(id, name, email, updates) VALUES(?, ?, ?, ?);"

	_, err = db.database.Exec(insert, p.id, p.name, p.email, updates)
	if err != nil {
		log.Fatal(err)
	}
}

func (db Database) ListPeople(w http.ResponseWriter) {
  rows, err := db.database.Query("SELECT id, name, email, updates FROM people;")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Fprintf(w, "%-5s %-15s %-20s %-5s\n", "ID", "Name", "Email", "Updates")
  fmt.Printf(strings.Repeat("-", 66) + "\n")

  for rows.Next() {
    var p Person
    rows.Scan(&p.id, &p.name, &p.email, &p.updates)
    fmt.Fprintf(w, "%-5d %-15s %-20s %-5d\n", p.id, p.name, p.email, p.updates)
  }
}

func (db Database) Close() {
	db.database.Close()
}
