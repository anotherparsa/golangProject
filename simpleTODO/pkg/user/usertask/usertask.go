package usertask

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
)

func CreateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		//checking if the request mothod is POST or not
		if r.Method == "POST" {
			//getting form input values
			r.ParseForm()
			//getting user's user_Id
			_, _, author := session.WhoIsThis(cookie.Value)
			//creating a task record in tasks table
			query, arguments := databasetools.QuerryMaker("insert", []string{"author", "priority", "category", "title", "description", "finished"}, "tasks", [][]string{}, [][]string{{"author", author}, {"priority", r.Form.Get("priority")}, {"category", r.Form.Get("category")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"finished", "unfinished"}})
			CreateTask(query, arguments)
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusMethodNotAllowed)
		}
	} else {
		http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
	}
}

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
		err := rows.Scan(&task.Id, &task.Author, &task.Priority, &task.Category, &task.Title, &task.Description, &task.Finished)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func UpdateTaskPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
	} else {
		_, _, userId := session.WhoIsThis(cookie.Value)
		taskID := strings.TrimPrefix(r.URL.Path, "/tasks/edittask/")
		Query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "finished"}, "tasks", [][]string{{"id", taskID}}, [][]string{})
		task := ReadTask(Query, arguments)
		template, err := template.ParseFiles("../../pkg/user/usertask/template/useredittask.html")
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

func UpdateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		Query, arguments := databasetools.QuerryMaker("update", []string{"priority", "title", "description", "category"}, "tasks", [][]string{{"id", r.Form.Get("id")}}, [][]string{{"priority", r.Form.Get("priority")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"category", r.Form.Get("category")}})
		UpdateTask(Query, arguments)
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	} else {
		fmt.Println("Wrong method")
		http.Redirect(w, r, "/users/home", http.StatusMethodNotAllowed)
	}
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

func DeleteTaskProcessor(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/tasks/deletetask/")
	query, arguments := databasetools.QuerryMaker("delete", []string{"id"}, "tasks", [][]string{{"id", taskID}}, [][]string{})
	DeleteTask(query, arguments)
	http.Redirect(w, r, "/users/home", http.StatusSeeOther)
}

func DeleteTask(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
}
