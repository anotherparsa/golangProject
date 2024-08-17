package task

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

//CRUD
//Create
func CreateTask(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		r.ParseForm()
		query, arguments := databasetools.QuerryMaker("insert", []string{"author", "priority", "title", "description", "isDone"}, "tasks", map[string]string{}, [][]string{{"author", r.Form.Get("author")}, {"priority", r.Form.Get("priority")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"isDone", "0"}})
		safequery, err := databasetools.DataBase.Prepare(query)
		if err != nil {
			fmt.Println(err)
		}
		_, err = safequery.Exec(arguments...)
		if err == nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

//Read
func ReadTask(db *sql.DB, query string, arguments []interface{}) []models.Task {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	tasks := []models.Task{}
	for rows.Next() {
		task := models.Task{}

		err := rows.Scan(&task.Id, &task.Author, &task.Priority, &task.Title, &task.Description, &task.IsDone)
		if err != nil {
			fmt.Println(err)
		}

		tasks = append(tasks, task)
	}
	return tasks
}

//Update
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query, arguments := databasetools.QuerryMaker("update", []string{"priority", "title", "description"}, "tasks", map[string]string{"id": r.Form.Get("id")}, [][]string{{"priority", r.Form.Get("priority")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}})
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//Delete
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/deletetask/")
	query, arguments := databasetools.QuerryMaker("delete", []string{"id"}, "tasks", map[string]string{"id": taskID}, [][]string{})
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
