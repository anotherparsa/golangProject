package adminusers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user/useruser"
)

type datatosend struct {
	CSRFT string
}

func AdminUsersManagementPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//checking if session_id exists or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		_, _, _, rule, _ := session.WhoIsThis(cookie.Value)
		//checking if the logged user is admin
		if rule == "admin" {
			csrft := tools.GenerateUUID()
			//setting csrft cookie
			http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
			datatosend := datatosend{CSRFT: csrft}
			template, _ := template.ParseFiles("../../pkg/admin/adminusers/template/adminusersmanagement.html")
			template.Execute(w, datatosend)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {

		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

func AdminUsersManagementProcess(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//checking if session_id exists or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		_, _, _, rule, _ := session.WhoIsThis(cookie.Value)
		//checking if the logged user is admin
		if rule == "admin" {
			generatedCSRFT, err := r.Cookie("adminupdateusercsrft")
			//checkinf if the csrft cookie exists
			if err == nil && generatedCSRFT != nil {
				r.ParseForm()
				//checking if the sent csrft is the same as generated one.
				if generatedCSRFT.Value == r.Form.Get("csrft") {
					//checking if the request method is equal to POST or not
					if r.Method == "POST" {
						//getting form input values:
						targetUsername := r.Form.Get("targetusername")
						operation := r.Form.Get("operation")
						//checking what to do with the target user
						if operation == "tempsuspend" {
							//we are going to suspend the user
							query, arguments := databasetools.QuerryMaker("update", []string{"suspended"}, "users", [][]string{{"username", targetUsername}}, [][]string{{"suspended", "yes"}})
							useruser.UpdateUser(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)

						} else if operation == "deleteuser" {
							//we are going to delete the user
							//getting the user
							query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{{"username", targetUsername}}, [][]string{})
							targetUser := useruser.ReadUser(query, arguments)
							//deleting user from users table
							query, arguments = databasetools.QuerryMaker("delete", []string{}, "users", [][]string{{"username", targetUser[0].Username}}, [][]string{})
							useruser.DeleteUserInfo(query, arguments)
							//deleting user's task from tasks table
							query, arguments = databasetools.QuerryMaker("delete", []string{}, "tasks", [][]string{{"author", targetUser[0].UserId}}, [][]string{})
							useruser.DeleteUserInfo(query, arguments)
							//deleting user's sessions from sessions table
							query, arguments = databasetools.QuerryMaker("delete", []string{}, "sessions", [][]string{{"userId", targetUser[0].UserId}}, [][]string{})
							useruser.DeleteUserInfo(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
							
						} else if operation == "promotetoadmin" {
							//we are going to promote the user to admin rule.
							query, arguments := databasetools.QuerryMaker("update", []string{"suspended"}, "users", [][]string{{"username", targetUsername}}, [][]string{{"rule", "admin"}})
							useruser.UpdateUser(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)

						} else if operation == "untempsuspend" {
							//we are going to unsuspend the user
							query, arguments := databasetools.QuerryMaker("update", []string{"suspended"}, "users", [][]string{{"username", targetUsername}}, [][]string{{"suspended", "no"}})
							useruser.UpdateUser(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)

						} else if operation == "unpromotetoadmin" {
							//we are going to demote the user to admin rule.
							query, arguments := databasetools.QuerryMaker("update", []string{"suspended"}, "users", [][]string{{"username", targetUsername}}, [][]string{{"rule", "user"}})
							useruser.UpdateUser(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)

						} else {
							http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
						http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
					http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
				http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "adminupdateusercsrft", MaxAge: -1, Path: "/"})
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func GetTotalUsers(query string, arguments []interface{}) string {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	counter := 0
	for rows.Next() {
		counter++
	}
	return strconv.Itoa(counter)
}
