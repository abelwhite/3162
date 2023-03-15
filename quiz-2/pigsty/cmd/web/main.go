package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)



// Dependency injection
type application struct {
	// pigs models.PigsModel
}

func main() {

	//Create a flag for specifing the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("PIGSTYDB_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	//create an instance to the application pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}

	//create a new instance of the application type
	app := &application{
		//pigs: models.PigsModel{DB: db}, //has to be of connection pool
	}

	defer db.Close() //if we dont close the application loop we have memory leak
	log.Println("Database connection pool established")

	//create a customized server
	srv := &http.Server{ //web server is listening for requests and send to router
		Addr:    *addr,
		Handler: app.routes(), //we created this routes server

	}

	//start our server
	log.Printf("Starting Server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err) //should never be reached

}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn) //check if dsn work
	if err != nil {
		return nil, err
	}
	//use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() //defer helps us not to use it in every if

	//lets ping the db
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
