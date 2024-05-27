package main

import (
	"log"

	"github.com/AndreBalDm/go_final_project/api"
	"github.com/AndreBalDm/go_final_project/db"
	"github.com/AndreBalDm/go_final_project/env"
)

func main() {
	env.SetFlagParams()
	err := db.DbExistance()
	if err != nil {
		log.Println("Err connect to DB:", err)
		return
	}
	api.StartWebServer()
}
