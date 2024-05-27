package api

import (
	"net/http"
	"time"

	"github.com/AndreBalDm/go_final_project/servicetask"
)

func (srv Server) GetTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]servicetask.Task{}
	var err error
	var tr servicetask.TaskResp

	searchString := r.FormValue("search")
	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}
	srv.Server.Response(tasks, w)
	switch searchString {
	//if no search take all records
	case "":
		tasks, tr, err = srv.Server.SrvService.GetAll()
		checkErr(err)

	//if search, then select the search bar
	default:
		searchDate, errParse := time.Parse("01.01.2001", searchString)
		//if search date
		if errParse == nil {
			tasks, err = srv.Server.SrvService.GetSearchDate(searchDate)
			checkErr(err)
			//if search no date
		} else {
			tasks, err = srv.Server.SrvService.GetSearch(searchString)
			checkErr(err)
		}
	}
}
