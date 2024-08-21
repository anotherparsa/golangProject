package models

//defining model structure for users, tasks, sessions.

type User struct {
	ID          int
	UserId      string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Rule        string
}

type Task struct {
	Id          int
	Author      string
	Priority    string
	Category    string
	Title       string
	Description string
	Finished    string
}

type Session struct {
	Id        int
	SessionId string
	UserId    string
}

type Messages struct {
	Id          int
	Author      string
	Priority    string
	Category    string
	Title       string
	Description string
	Finished    string
}

type Static struct {
	Id            int
	Totalusers    string
	Totaltasks    string
	Totalmessages string
}
