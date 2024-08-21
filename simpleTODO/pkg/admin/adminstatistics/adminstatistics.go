package adminstatistics

import (
	"todoproject/pkg/admin/adminmessages"
	"todoproject/pkg/admin/admintasks"
	"todoproject/pkg/admin/adminusers"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

func InitializeStaticsProcess() models.Static {
	query, arguments := databasetools.QuerryMaker("select", []string{"id"}, "users", [][]string{}, [][]string{})
	totalusers := adminusers.GetTotalUsers(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "messages", [][]string{}, [][]string{})
	totalmessages := adminmessages.GetTotalMessages(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "tasks", [][]string{}, [][]string{})
	totaltasks := admintasks.GetTotalTasks(query, arguments)
	data := models.Static{Totalusers: totalusers, Totaltasks: totaltasks, Totalmessages: totalmessages}
	return data
}
