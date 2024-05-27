package serverservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndreBalDm/go_final_project/servicetask"
)

type ServerService struct {
	SrvService servicetask.TaskStore
}

func NewServerService(SrvService servicetask.TaskStore) ServerService {
	return ServerService{SrvService: SrvService}
}

// check valid request
func (ss ServerService) ReqValidate(t *servicetask.Task) (servicetask.TaskResp, error) {
	//check all date & title in task valid
	var tr servicetask.TaskResp
	err := checkFieldsTask(t)
	if err != nil {
		tr.Err = "err format date or title"
	}
	return tr, nil
}

func (ss ServerService) Response(t any, w http.ResponseWriter) {
	resp, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (ss ServerService) RequestUpd(r *http.Request) (servicetask.Task, error) {
	var buf bytes.Buffer
	var task servicetask.Task
	//take data from web
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return servicetask.Task{}, err
	}
	//transfer date into struct task
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		return servicetask.Task{}, err
	}
	return task, nil
}

func (ss ServerService) RequestId(r *http.Request) int {
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	return idInt

}
