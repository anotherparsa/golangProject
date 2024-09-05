package usermessages

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
)

type datatosend struct {
	CSRFT string
}

func UserMessagePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//checking if session id exist or not, that means if the user is logged or not
	if err == nil && cookie != nil {
		csrft := tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		//parsing and rendering the sending message page.
		datatosend := datatosend{CSRFT: csrft}
		template, _ := template.ParseFiles("../../pkg/user/usermessages/template/usermessage.html")
		template.Execute(w, datatosend)
	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

func CreateUserMessageProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		generatedCSRFT, err := r.Cookie("createmessagecsrft")
		//checking if the csrft cookie exist or not
		if err == nil && generatedCSRFT != nil {
			r.ParseForm()
			//checking if the form csrft is the same as generated csrft at server
			if generatedCSRFT.Value == r.Form.Get("csrft") {
				//checking if the request method is POST or not
				if r.Method == "POST" {
					//getting logged user's username
					username, _, _, _, _ := session.WhoIsThis(cookie.Value)
					//getting users input values
					priority := r.Form.Get("priority")
					category := r.Form.Get("category")
					title := r.Form.Get("title")
					description := r.Form.Get("description")
					//form input validation
					if databasetools.ValidateTaskOrMessageInfoFormInputs("priority", priority) {
						if databasetools.ValidateTaskOrMessageInfoFormInputs("category", category) {
							if databasetools.ValidateTaskOrMessageInfoFormInputs("title", title) {
								if databasetools.ValidateTaskOrMessageInfoFormInputs("description", description) {
									//creating a message record in messages table
									query, arguments := databasetools.QueryMaker("insert", []string{"author", "priority", "category", "title", "description", "status"}, "messages", [][]string{}, [][]string{{"author", username}, {"priority", priority}, {"category", category}, {"title", title}, {"description", description}, {"status", "unfinished"}})
									CreateMessage(query, arguments)
									http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/home", http.StatusSeeOther)
								} else {
									http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
								http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
						http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
					http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
				http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/users/messages", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "createmessagecsrft", MaxAge: -1, Path: "/"})
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

func CreateMessage(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
}
