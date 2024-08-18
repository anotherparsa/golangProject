package models

type User struct {
	ID          int
	UserId      string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

type Task struct {
	Id          int
	Author      string
	Priority    string
	Category    string
	Title       string
	Description string
	IsDone      string
}

type Session struct {
	Id        int
	SessionId string
	UserId    string
}
