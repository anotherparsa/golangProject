package adminstatistics

import (
	"fmt"
	"strconv"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

func InitializeStaticsProcess() models.Static {
	query, arguments := databasetools.QuerryMaker("select", []string{"id"}, "users", [][]string{}, [][]string{})
	totalusers := GetTotal(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "messages", [][]string{}, [][]string{})
	totalmessagesnumber := GetTotal(query, arguments)
	query, arguments = databasetools.QuerryMaker("select", []string{"id"}, "tasks", [][]string{}, [][]string{})
	totaltasks := GetTotal(query, arguments)
	data := models.Static{Totalusers: totalusers, Totaltasks: totaltasks, Totalmessages: totalmessagesnumber}
	return data
}

func GetTotal(query string, arguments []interface{}) string {
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
