package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"snowlabs/chrono/components"
	"snowlabs/chrono/helpers"
	"snowlabs/chrono/models"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "chrono",
		AllowNativePasswords: true,
	}

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
	rows, err := db.Query("SELECT * FROM events ORDER BY start DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Start, &event.End, &event.Platform)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		event.Start = helpers.UnixTime(event.Start)
		event.End = helpers.UnixTime(event.End)
		event.Name = helpers.RemoveUnderscores(event.Name)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	components.Root(events).Render(r.Context(), w)
}
