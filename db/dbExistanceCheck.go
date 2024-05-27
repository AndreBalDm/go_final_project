package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AndreBalDm/go_final_project/env"
	_ "modernc.org/sqlite"
)

func DbExistance() error {
	dbFile := env.DbName()
	_, err := os.Stat(dbFile)
	if err != nil {
		log.Println("Create new DB with table scheduler")
		err = dbCreate(dbFile)
		if err != nil {
			return fmt.Errorf("err create new DB: %w", err)
		}
	}
	log.Println("Connect to DB")
	return nil
}

func dbCreate(dbFile string) error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("Err connect to DB: %w", err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT,	
											date CHAR(8), 
											title VARCHAR(256) NOT NULL DEFAULT "", 
											comment VARCHAR(256) ,
											repeat VARCHAR(256))`)
	if err != nil {
		return fmt.Errorf("err create table in DB: %w", err)
	}
	_, err = db.Exec("CREATE INDEX dateindex ON scheduler (date)")
	if err != nil {
		return fmt.Errorf("err create dateindex in DB: %w", err)
	}
	return nil
}
