package usertask

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
)

type datatosend struct {
	CSRFT string
	Task  models.Task
}

func CreateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		generatedCSRFT, err := r.Cookie("homecsrft")
		//checking if the csrft cookie exist or not
		if err == nil && generatedCSRFT != nil {
			r.ParseForm()
			//checking if the form csrft is the same as generated csrft at server
			if generatedCSRFT.Value == r.Form.Get("csrft") {
				//checking if the request mothod is POST or not
				if r.Method == "POST" {
					//getting user's user_Id
					_, _, author := session.WhoIsThis(cookie.Value)
					//getting users input values
					priority := r.Form.Get("priority")
					category := r.Form.Get("category")
					title := r.Form.Get("title")
					description := r.Form.Get("description")
					//validating users input on tasks forms
					if tools.ValidateTaskOrMessageInfoFormInputs("priority", priority) {
						if tools.ValidateTaskOrMessageInfoFormInputs("category", category) {
							if tools.ValidateTaskOrMessageInfoFormInputs("title", title) {
								if tools.ValidateTaskOrMessageInfoFormInputs("description", description) {
									//creating a task record in tasks table
									query, arguments := databasetools.QuerryMaker("insert", []string{"author", "priority", "category", "title", "description", "status"}, "tasks", [][]string{}, [][]string{{"author", author}, {"priority", r.Form.Get("priority")}, {"category", r.Form.Get("category")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"status", "unfinished"}})
									CreateTask(query, arguments)
									http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/home", http.StatusSeeOther)
								} else {
									http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/home", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
								http.Redirect(w, r, "/users/home", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/users/home", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
						http.Redirect(w, r, "/users/home", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
				http.Redirect(w, r, "/users/home", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
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
		err := rows.Scan(&task.Id, &task.Author, &task.Priority, &task.Category, &task.Title, &task.Description, &task.Status)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func UpdateTaskPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session id exist or not, that means if the user is logged in
	if err == nil || cookie != nil {
		csrft := tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		//getting logged user userId
		_, _, userId := session.WhoIsThis(cookie.Value)
		taskId := strings.TrimPrefix(r.URL.Path, "/tasks/edittask/")
		//getting task to edit
		Query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "status"}, "tasks", [][]string{{"id", taskId}, {"author", userId}}, [][]string{})
		task := ReadTask(Query, arguments)
		//checking if it had any result or not
		if len(task) == 1 {
			datatosend := datatosend{CSRFT: csrft, Task: task[0]}
			//rendering edit page with the target task
			template, _ := template.ParseFiles("../../pkg/user/usertask/template/useredittask.html")
			template.Execute(w, datatosend)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

func UpdateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		generatedCSRFT, err := r.Cookie("updatetaskcsrft")
		_, _, loggedUser := session.WhoIsThis(cookie.Value)
		//checking if the csrft cookie exists
		if err == nil && generatedCSRFT != nil {
			r.ParseForm()
			//checking if the sent csrft is the same as generated one
			if generatedCSRFT.Value == r.Form.Get("csrft") {
				//checking if the request method is equal to POST
				if r.Method == "POST" {
					//getting form inputs value
					priority := r.Form.Get("priority")
					category := r.Form.Get("category")
					title := r.Form.Get("title")
					description := r.Form.Get("description")
					if tools.ValidateTaskOrMessageInfoFormInputs("priority", priority) {
						if tools.ValidateTaskOrMessageInfoFormInputs("category", category) {
							if tools.ValidateTaskOrMessageInfoFormInputs("title", title) {
								if tools.ValidateTaskOrMessageInfoFormInputs("description", description) {
									//getting task to edit
									Query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "status"}, "tasks", [][]string{{"id", r.Form.Get("id")}, {"author", loggedUser}}, [][]string{})
									task := ReadTask(Query, arguments)
									//checking if it had any result or not
									if len(task) == 1 {
										Query, arguments := databasetools.QuerryMaker("update", []string{"priority", "title", "description", "category"}, "tasks", [][]string{{"id", r.Form.Get("id")}, {"author", loggedUser}}, [][]string{{"priority", priority}, {"title", title}, {"description", description}, {"category", category}})
										UpdateTask(Query, arguments)
										http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
										http.Redirect(w, r, "/users/home", http.StatusSeeOther)
									} else {
										http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
										http.Redirect(w, r, "/users/home", http.StatusSeeOther)
									}
								} else {
									http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/home", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
								http.Redirect(w, r, "/users/home", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/users/home", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
						http.Redirect(w, r, "/users/home", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
				http.Redirect(w, r, "/users/home", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "updatetaskcsrft", MaxAge: -1, Path: "/"})
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

//applying in the database
func UpdateTask(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
}

//Delete
//Deleting processor
func DeleteTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		//getting the userId of the logged user
		_, _, loggedUser := session.WhoIsThis(cookie.Value)
		//getting the task id from url
		taskID := strings.TrimPrefix(r.URL.Path, "/tasks/deletetask/")
		//getting the task with that id and that userId as author
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "status"}, "tasks", [][]string{{"id", taskID}, {"author", loggedUser}}, [][]string{})
		task := ReadTask(query, arguments)
		//checking if there was a task to meet those conditions.
		if len(task) == 1 {
			//deleting task
			query, arguments = databasetools.QuerryMaker("delete", []string{"id"}, "tasks", [][]string{{"id", taskID}}, [][]string{})
			DeleteTask(query, arguments)
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

//applying in the database
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
