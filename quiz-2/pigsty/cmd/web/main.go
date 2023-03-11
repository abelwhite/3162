package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config settings
type config struct {
	port int
	env  string
}

// Dependency injection
type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	//Create a flag for specifing the port number when starting the server
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	//dsn := flag.String("dsn", os.Getenv("PIGSTYDB_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//instance of application
	app := &application{
		config: cfg,
		logger: logger,
	}

	//create a new server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/viewpig", app.Viewpig)

	//create our http server
	srv := &http.Server{ //web server is listening for requests and send to router
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//start our server
	logger.Printf("Starting %s server om %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
