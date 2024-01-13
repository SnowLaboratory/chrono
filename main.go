package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"snowlabs/chrono/helpers"
	"strconv"
	"text/template"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Event struct {
	ID       int
	Name     string
	Start    string
	End      string
	Platform string
}

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "chrono",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	r := mux.NewRouter()
	r.HandleFunc("/", getAllEvents).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./tmpl/index.html"))

	rows, err := db.Query("SELECT * FROM events ORDER BY start DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Start, &event.End, &event.Platform)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		i, err := strconv.ParseInt(event.Start, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		tm := time.Unix(i, 0)
		event.Start = tm.Format("2006-01-02 15:04:05 MST")
		event.Name = helpers.RemoveUnderscores(event.Name)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, events)
}
