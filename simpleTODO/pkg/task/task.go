package task

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
)

type Task struct {
	Id          string
	Priority    string
	Title       string
	Description string
	IsDone      string
}

func CreateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		r.ParseForm()
		author := databasetools.WhoIsThis(databasetools.DB, cookie.Value)
		priority := r.Form.Get("priority")
		title := r.Form.Get("title")
		description := r.Form.Get("description")
		databasetools.CreateTasks(databasetools.DB, author, priority, title, description)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func GetUsersTask(db *sql.DB, username string) ([]Task, error) {
	rows, err := db.Query("SELECT id, priority, title, description, isDone FROM tasks WHERE author = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Priority, &task.Title, &task.Description, &task.IsDone); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetUserTaskByTaskID(db *sql.DB, id string) Task {
	rows, err := db.Query("SELECT id, priority, title, description, isDone FROM tasks WHERE id=?", id)
	var t Task
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.Id, &t.Priority, &t.Title, &t.Description, &t.IsDone); err != nil {
			fmt.Println(err)
		}
	}
	return t

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/deletetask/")
	databasetools.DeleteTask(databasetools.DB, taskID)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	newTitle := r.Form.Get("title")
	newDescription := r.Form.Get("description")
	newPriority := r.Form.Get("priority")
	databasetools.EditTask(databasetools.DB, id, newTitle, newDescription, newPriority)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func EditTaskPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../../pkg/task/template/edittask.html")
	taskID := strings.TrimPrefix(r.URL.Path, "/edittask/")
	task := GetUserTaskByTaskID(databasetools.DB, taskID)
	t.Execute(w, task)
}
