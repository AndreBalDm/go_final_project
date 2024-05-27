package servicetask

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/AndreBalDm/go_final_project/nextdate"
	"github.com/AndreBalDm/go_final_project/params"
)

const limit = 15

// initializing DB
func NewTaskStore(Db *sql.DB) TaskStore {
	return TaskStore{Db: Db}
}

// add to DB & back number task & err
func (ts TaskStore) Add(t *Task) (TaskResp, error) {
	var tr TaskResp
	//write struct Task in DB
	res, err := ts.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat))
	if err != nil {
		return TaskResp{}, fmt.Errorf("err add task DB: %w", err)
	}

	//get ID last task
	lastID, err := res.LastInsertId()
	if err != nil {
		return tr, fmt.Errorf("err get last ID: %w", err)
	}

	tr.Id = strconv.Itoa(int(lastID))
	return tr, nil
}

// get task ID
func (ts TaskStore) GetOneTask(id int) (Task, TaskResp, error) {
	var task Task
	var tr TaskResp
	err := ts.Db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		tr.Err = "err no this ID"
		return Task{}, tr, fmt.Errorf("err read data id: %w", err)
	}
	return task, tr, nil
}

// get task if search no date
func (ts TaskStore) GetSearch(searchString string) (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	rows, err := ts.Db.Query("SELECT * FROM scheduler WHERE title LIKE :searchString OR comment LIKE :searchString ORDER BY date LIMIT :limit",
		sql.Named("searchString", "%"+searchString+"%"),
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("err request DB: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, fmt.Errorf("err parsing after read from DB: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, nil
}

// et task if search date
func (ts TaskStore) GetSearchDate(searchDate time.Time) (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	rows, err := ts.Db.Query("SELECT * FROM scheduler WHERE date = :searchString LIMIT :limit",
		sql.Named("searchString", searchDate.Format(params.DFormat)),
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("err request DB: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, fmt.Errorf("err parsing after read from DB: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, nil
}

// get all task from DB
func (ts TaskStore) GetAll() (map[string][]Task, TaskResp, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var tr TaskResp
	var task Task
	rows, err := ts.Db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
		sql.Named("limit", limit))

	if err != nil {
		tr.Err = "err request DB"
		return tasks, tr, fmt.Errorf("err request DB: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			tr.Err = "err parsing after read from DB"
			return tasks, tr, fmt.Errorf("err parsing after read from DB: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, tr, nil
}

// delete task
func (ts TaskStore) Delete(id int) (TaskResp, error) {
	var tr = TaskResp{}
	var checkedID string

	err := ts.Db.QueryRow("SELECT id FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&checkedID)
	if err != nil {
		tr.Err = "err no this ID"
		return tr, fmt.Errorf("err read data id: %w", err)
	}
	_, err = ts.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return TaskResp{}, fmt.Errorf("err update task in DB: %w", err)
	}
	return tr, nil
}

// done task
func (ts TaskStore) Done(id int) (TaskResp, error) {
	var tr TaskResp
	var task Task

	err := ts.Db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return TaskResp{}, fmt.Errorf("err read data id: %w", err)
	}

	//check ID task
	if len(task.Id) == 0 {
		tr.Err = "err no this ID"
	}

	if task.Repeat == "" {
		_, err = ts.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
		if err != nil {
			return TaskResp{}, fmt.Errorf("err update task in DB: %w", err)
		}
	} else {
		newDate, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			return TaskResp{}, fmt.Errorf("err new date: %w", err)
		}
		_, err = ts.Db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
			sql.Named("date", newDate),
			sql.Named("id", task.Id))
		if err != nil {
			return TaskResp{}, fmt.Errorf("err update task in DB: %w", err)
		}
	}
	return tr, nil
}

// updates Task
func (ts TaskStore) Update(t Task) (TaskResp, error) {
	var tr TaskResp
	//updates structы task in DB
	result, err := ts.Db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.Id))
	if err != nil {
		return TaskResp{}, fmt.Errorf("err update task in DB: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return TaskResp{}, fmt.Errorf("err update task in DB: %w", err)
	}
	if rowsAffected == 0 {
		tr.Err = "task no find"
	}

	return tr, nil
}
