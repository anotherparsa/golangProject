package statics

import (
	"todoproject/pkg/admin/messages"
	"todoproject/pkg/admin/tasks"
	"todoproject/pkg/admin/users"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

func InitializeStaticsProcess() models.Static {
	query, arguments := databasetools.QuerryMaker("select", []string{"id"}, "users", [][]string{}, [][]string{})
	totalusers := users.GetTotalUsers(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "messages", [][]string{}, [][]string{})
	totalmessages := messages.GetTotalMessages(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "tasks", [][]string{}, [][]string{})
	totaltasks := tasks.GetTotalTasks(query, arguments)
	data := models.Static{Totalusers: totalusers, Totaltasks: totaltasks, Totalmessages: totalmessages}
	return data
}
