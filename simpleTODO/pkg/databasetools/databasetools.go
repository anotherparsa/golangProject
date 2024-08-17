package databasetools

import (
	"database/sql"
	"fmt"
)

const (
	username = "testuser"
	password = "testpass"
	hostname = "localhost"
	port     = "3306"
	database = "users"
)

var DataBase *sql.DB

func ConnectToDatabase() (*sql.DB, error) {
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	//testuser:testpass@tcp(localhost:3306)/users
	return sql.Open("mysql", DataSourceName)
}

func CreateDatabase() {
	DataBase, _ = ConnectToDatabase()
}

//this way the code is vulnerable to sql injection
//fix that
func QuerryMaker(operation string, columns []string, table string, conditions map[string]string, values map[string]string) (string, []interface{}) {
	var args []interface{} // to hold the actual values
	var query string

	if operation == "select" {
		query = "SELECT "
		for i := 0; i < len(columns); i++ {
			if i < len(columns)-1 {
				query += fmt.Sprintf("%v, ", columns[i])
			} else {
				query += fmt.Sprintf("%v ", columns[i])
			}
		}
		query += fmt.Sprintf("FROM %v", table)

		if len(conditions) != 0 {
			query += " WHERE "
			first := true
			for columnName, value := range conditions {
				if !first {
					query += " AND "
				}
				query += fmt.Sprintf("%v = ?", columnName)
				args = append(args, value) // collect the condition value
				first = false
			}
		}
		return query, args

	} else if operation == "update" {
		query = "UPDATE " + table + " SET "
		first := true
		for column, value := range values {
			if !first {
				query += ", "
			}
			query += fmt.Sprintf("%v = ?", column)
			args = append(args, value) // collect the update value
			first = false
		}

		if len(conditions) != 0 {
			query += " WHERE "
			first = true
			for columnName, value := range conditions {
				if !first {
					query += " AND "
				}
				query += fmt.Sprintf("%v = ?", columnName)
				args = append(args, value) // collect the condition value
				first = false
			}
		}
		return query, args

	} else if operation == "insert" {
		query = fmt.Sprintf("INSERT INTO %v (", table)
		for i := 0; i < len(columns); i++ {
			if i < len(columns)-1 {
				query += fmt.Sprintf("%v, ", columns[i])
			} else {
				query += fmt.Sprintf("%v ", columns[i])
			}
		}
		query += ") VALUES ("

		for i := 0; i < len(values); i++ {
			if i < len(values)-1 {
				query += "?, "
			} else {
				query += "? "
			}
			// Collecting values to args slice
			for _, value := range values {
				args = append(args, value)
			}
		}
		query += ")"
		return query, args
	}

	return "invalid operation", nil
}
