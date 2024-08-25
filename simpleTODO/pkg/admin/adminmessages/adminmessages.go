package adminmessages

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

func AdminMessagesPageHandler(w http.ResponseWriter, r *http.Request) {
	query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "status"}, "messages", [][]string{}, [][]string{})
	messages := GetMessages(query, arguments)
	template, _ := template.ParseFiles("../../pkg/admin/adminmessages/template/adminmessages.html")
	template.Execute(w, messages)
}

func GetMessages(query string, arguments []interface{}) []models.Messages {
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
	for rows.Next() {
		message := models.Messages{}
		err := rows.Scan(&message.Id, &message.Author, &message.Priority, &message.Title, &message.Description, &message.Description, &message.Status)
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, message)
	}
	return messages
}
