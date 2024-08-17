package databasetools

import (
	"database/sql"
	"fmt"
	"regexp"
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
	fmt.Println(DataBase)
}

func isValidIdentifier(identifier string) bool {
	// This regex allows alphanumeric characters and underscores
	// You may want to adjust this based on your database's naming conventions
	validIdentifier := regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	return validIdentifier.MatchString(identifier)
}

func QuerryMaker(operation string, columns []string, table string, conditions map[string]string, values [][]string) (string, []interface{}) {
	var args []interface{}
	var query string

	if !isValidIdentifier(table) {
		return "invalid table name", nil
	}

	for _, col := range columns {
		if !isValidIdentifier(col) {
			return "invalid column name", nil
		}
	}

	if operation == "select" {
		query = "SELECT "
		for i, col := range columns {
			if i < len(columns)-1 {
				query += fmt.Sprintf("%s, ", col)
			} else {
				query += fmt.Sprintf("%s ", col)
			}
		}
		query += fmt.Sprintf("FROM %s", table)

		if len(conditions) != 0 {
			query += " WHERE "
			first := true
			for columnName, value := range conditions {
				if !first {
					query += " AND "
				}
				query += fmt.Sprintf("%s = ?", columnName)
				args = append(args, value)
				first = false
			}
		}
		return query, args

	} else if operation == "update" {
		query = "UPDATE " + table + " SET "
		first := true
		for _, valuePair := range values {
			if len(valuePair) != 2 {
				return "invalid value pair", nil
			}
			column := valuePair[0]
			value := valuePair[1]

			if !isValidIdentifier(column) {
				return "invalid column name", nil
			}

			if !first {
				query += ", "
			}
			query += fmt.Sprintf("%s = ?", column)
			args = append(args, value)
			first = false
		}

		if len(conditions) != 0 {
			query += " WHERE "
			first = true
			for columnName, value := range conditions {
				if !first {
					query += " AND "
				}
				query += fmt.Sprintf("%s = ?", columnName)
				args = append(args, value)
				first = false
			}
		}
		return query, args

	} else if operation == "insert" {
		query = fmt.Sprintf("INSERT INTO %s (", table)
		for i, col := range columns {
			if i < len(columns)-1 {
				query += fmt.Sprintf("%s, ", col)
			} else {
				query += fmt.Sprintf("%s ", col)
			}
		}
		query += ") VALUES ("
		fmt.Println(values)
		for i := 0; i < len(values); i++ {
			if len(values[i]) != 2 {
				return "invalid value pair", nil
			}
			if i < len(values)-1 {
				query += "?, "
			} else {
				query += "? "
			}
			fmt.Println(values[i][1])
			args = append(args, values[i][1]) // Add the value part to args
		}
		query += ")"

		fmt.Println("query making was ok")
		return query, args
	}

	return "invalid operation", nil
}
