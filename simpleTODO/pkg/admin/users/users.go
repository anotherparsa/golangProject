package users

import (
	"fmt"
	"strconv"
	"todoproject/pkg/databasetools"
)

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
