package api

import (
	"log"
	"net/http"
	"time"

	"github.com/AndreBalDm/go_final_project/nextdate"
	"github.com/AndreBalDm/go_final_project/params"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	//get the parameters from the request and convert NOW to the time format
	nowInString := r.FormValue("now")
	now, err := time.Parse(params.DFormat, nowInString)
	if err != nil {
		log.Println("date format err:", err)
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	// get new date
	s, err := nextdate.NextDate(now, date, repeat)
	if err != nil {
		log.Println("date transfer err:", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}
