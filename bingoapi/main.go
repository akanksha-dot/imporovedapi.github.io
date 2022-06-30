package main

import (
	"context"
	"database/sql"

	"encoding/json"

	_ "fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var db *sql.DB

var ctx context.Context

func getMYSQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:suspiciouscoder19@/country")
	_ = err

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	return db
}

func getcountrylist(w http.ResponseWriter, r *http.Request) {
	db = getMYSQLDB()
	rows, err := db.QueryContext(ctx, "SELECT * FROM list_country")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(err)
		// fmt.Printf("%s is %d\n", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	/* we are creating r (ROUTER)*/
	r := mux.NewRouter()
	/* api function*/
	r.HandleFunc("/countries", getcountrylist).Methods("GET")
	http.ListenAndServe(":8080", r)
}
