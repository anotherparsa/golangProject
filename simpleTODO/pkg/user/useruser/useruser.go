package useruser

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
	User  models.User
}

func CreateUser(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadUser(query string, arguments []interface{}) []models.User {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	user := models.User{}
	users := []models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users
}

func UpdateUserPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//checking if session id exist or not, that means if the user is logged in or not.
	if err == nil && cookie != nil {
		csrft := tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		//getting logged user userId and id
		_, usersid, _, _, _ := session.WhoIsThis(cookie.Value)
		usersIdurl := strings.TrimPrefix(r.URL.Path, "/users/editaccount/")
		//checking if the id of the logged user is same as the id in url path
		if usersid == usersIdurl {
			Query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"id", usersid}}, [][]string{})
			user := ReadUser(Query, arguments)
			template, err := template.ParseFiles("../../pkg/user/useruser/template/useredituser.html")
			if err != nil {
				fmt.Println(err)
			}
			datatosend := datatosend{CSRFT: csrft, User: user[0]}
			template.Execute(w, datatosend)
		} else {
			http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}

func UpdateUserProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//checking if session id exist or not, that means if the user is logged in or not.
	if err == nil && cookie != nil {
		generatedCSRFT, err := r.Cookie("updateusercsrft")
		_, _, loggedUser, _, _ := session.WhoIsThis(cookie.Value)
		//checking if the csrft cookie exists
		if err == nil && generatedCSRFT != nil {
			r.ParseForm()
			if generatedCSRFT.Value == r.Form.Get("csrft") {
				//checking if the request method is equal to POST
				if r.Method == "POST" {
					username := r.Form.Get("username")
					currentpassword := r.Form.Get("password")
					newpassword := r.Form.Get("newpassword")
					firstName := r.Form.Get("firstName")
					lastName := r.Form.Get("lastName")
					email := r.Form.Get("email")
					phoneNumber := r.Form.Get("phoneNumber")
					//checking if forms input are valid or not
					if tools.ValidateUserInfoFormInputs("username", username) {
						if tools.ValidateUserInfoFormInputs("password", currentpassword) {
							if tools.ValidateUserInfoFormInputs("firstName", firstName) {
								if tools.ValidateUserInfoFormInputs("lastName", lastName) {
									if tools.ValidateUserInfoFormInputs("email", email) {
										if tools.ValidateUserInfoFormInputs("phoneNumber", phoneNumber) {
											//getting user to edit
											Query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"id", r.Form.Get("id")}, {"userId", loggedUser}, {"password", tools.HashThis(r.Form.Get("currentpassword"))}}, [][]string{})
											user := ReadUser(Query, arguments)
											//checking if it had any result or not
											if len(user) == 1 {
												//checkinf if the user entered a new password
												if len(newpassword) != 0 {
													if tools.ValidateTaskOrMessageInfoFormInputs("password", newpassword) {
														Query, arguments := databasetools.QuerryMaker("update", []string{"username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"userId", loggedUser}}, [][]string{{"username", r.Form.Get("username")}, {"password", tools.HashThis(r.Form.Get("newpassword"))}, {"firstName", r.Form.Get("FirstName")}, {"lastName", r.Form.Get("LastName")}, {"email", r.Form.Get("Email")}, {"phoneNumber", r.Form.Get("PhoneNumber")}})
														UpdateUser(Query, arguments)
														http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
														http.Redirect(w, r, "/users/home", http.StatusSeeOther)
													} else {
														http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
														http.Redirect(w, r, "/users/home", http.StatusSeeOther)
													}
												} else {
													//this means user didn't provide a new password
													Query, arguments := databasetools.QuerryMaker("update", []string{"username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"userId", loggedUser}}, [][]string{{"username", r.Form.Get("username")}, {"firstName", r.Form.Get("FirstName")}, {"lastName", r.Form.Get("LastName")}, {"email", r.Form.Get("Email")}, {"phoneNumber", r.Form.Get("PhoneNumber")}})
													UpdateUser(Query, arguments)
													http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
													http.Redirect(w, r, "/users/home", http.StatusSeeOther)
												}
											} else {
												http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
												http.Redirect(w, r, "/users/home", http.StatusSeeOther)
											}
										} else {
											http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
											http.Redirect(w, r, "/users/home", http.StatusSeeOther)
										}
									} else {
										http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
										http.Redirect(w, r, "/users/home", http.StatusSeeOther)
									}
								} else {
									http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
									http.Redirect(w, r, "/users/home", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
								http.Redirect(w, r, "/users/home", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
							http.Redirect(w, r, "/users/home", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
						http.Redirect(w, r, "/users/home", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
				http.Redirect(w, r, "/users/home", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "updateusercsrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func UpdateUser(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
}
