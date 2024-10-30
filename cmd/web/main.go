package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	envError := godotenv.Load()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	if envError != nil {
		errorLog.Fatal("Error loading .env file")
	}
	//define port command line flag
	defaultPort := os.Getenv("SNIPPETBOX_ADDR")
	addr := flag.String("addr", defaultPort, "HTTP network address")
	flag.Parse() //parse the flags so they can be used
	// define command line flag for MYSQL DSN string
	dsn := flag.String("dsn", os.Getenv("SNIPPETBOX_DB"), "MySQL data source name")
	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	} else {
		infoLog.Println("Database connection successful")
	}
	defer db.Close()
	// define new instance of app containing the dependecies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
