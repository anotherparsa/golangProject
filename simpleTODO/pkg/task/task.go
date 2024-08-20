package task

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
)

//CRUD operations for tasks
//Create
//processor
func CreateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		r.ParseForm()
		author := session.ReturnUsersUserID(cookie.Value)
		query, arguments := databasetools.QuerryMaker("insert", []string{"author", "priority", "title", "description", "isDone"}, "tasks", [][]string{}, [][]string{{"author", author}, {"priority", r.Form.Get("priority")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"isDone", "0"}})
		CreateTask(query, arguments)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

//apply in database
func CreateTask(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Task has been created")
}

//Read
//processor
func ReadTask(query string, arguments []interface{}) []models.Task {
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
//page handler
func UpdateTaskPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		_, _, userId := session.WhoIsThis(cookie.Value)

		taskID := strings.TrimPrefix(r.URL.Path, "/edittask/")
		Query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "title", "description", "isDone"}, "tasks", [][]string{{"id", taskID}}, [][]string{})
		task := ReadTask(Query, arguments)
		template, err := template.ParseFiles("../../pkg/task/template/edittask.html")
		if err != nil {
			fmt.Println(err)
		}
		if userId != task[0].Author {
			fmt.Fprintf(w, "You are not authorized")
		} else {
			template.Execute(w, task[0])
		}

	}

}

//processor
func UpdateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Query, arguments := databasetools.QuerryMaker("update", []string{"priority", "title", "description"}, "tasks", [][]string{{"id", r.Form.Get("id")}}, [][]string{{"priority", r.Form.Get("priority")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}})
	UpdateTask(Query, arguments)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func UpdateTask(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("task has been updated")
}

//Delete
//processor
func DeleteTaskProcessor(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/deletetask/")
	query, arguments := databasetools.QuerryMaker("delete", []string{"id"}, "tasks", [][]string{{"id", taskID}}, [][]string{})
	DeleteTask(query, arguments)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//apply in database
func DeleteTask(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Task has been deleted")
}
