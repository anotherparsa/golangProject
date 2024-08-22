package adminmessages

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

func AdminMessagesPageHandler(w http.ResponseWriter, r *http.Request) {
	query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "finished"}, "messages", [][]string{}, [][]string{})
	_, messages := GetTotalMessages(query, arguments)
	template, _ := template.ParseFiles("../../pkg/admin/adminmessages/template/adminmessages.html")
	template.Execute(w, messages)
}

func GetTotalMessages(query string, arguments []interface{}) (string, []models.Messages) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	messages := []models.Messages{}
	counter := 0
	for rows.Next() {
		message := models.Messages{}
		err := rows.Scan(&message.Id, &message.Author, &message.Priority, &message.Title, &message.Description, &message.Description, &message.Finished)
		if err != nil {
			fmt.Println(err)
		}
		counter++
		messages = append(messages, message)
	}
	return strconv.Itoa(counter), messages
}
