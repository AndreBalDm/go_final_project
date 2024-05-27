package env

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

// identify point
func SetFlagParams() {
	pass := flag.String("password", "", "Password for app")
	port := flag.String("port", "7540", "PORT for startweb server")
	dbPath := flag.String("dbpath", "", "Path for DB")
	flag.Parse()
	os.Setenv("TODO_PASSWORD", *pass)
	os.Setenv("TODO_PORT", *port)
	os.Setenv("TODO_DBFILE", *dbPath)
}

func SetPass() string {
	return os.Getenv("TODO_PASSWORD")
}

func SetPort() string {
	return os.Getenv("TODO_PORT")
}

func DbName() string {
	dbFile := "scheduler.db"
	envFile := os.Getenv("TODO_DBFILE")
	if len(envFile) > 0 {
		dbFile = filepath.Join(envFile, "scheduler.db")
	}
	log.Println("путь к БД:", dbFile)
	return dbFile
}
