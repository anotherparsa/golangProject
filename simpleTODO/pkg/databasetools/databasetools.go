package databasetools

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"
	"todoproject/pkg/models"
	"todoproject/pkg/tools"

	_ "github.com/go-sql-driver/mysql"
)

// Should not be a singleton
var DataBase *sql.DB

const (
	// should be read from env variable via os.Getenv or packages that do that
	username = "testuser"
	password = "testpass"
	hostname = "localhost"
	port     = "3306"
	database = "todo"
)

// opening a connection
func connect() (*sql.DB, error) {
	// you could use os.ExpandEnv to read everything from the env variable in one go
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	return sql.Open("mysql", dsn)
}

// caling the connnect function
func CreateDatabase() {
	// catch the error, this could really happen, and you want to know why it didn't work,
	// so at least log it, or return one level up
	DataBase, _ = connect()
}
func isValidIdentifier(identifier string) bool {
	validIdentifier := regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	return validIdentifier.MatchString(identifier)
}

// making queries for sql
// this seems like a bad idea that can have bad side effect that are hard to detect
func QueryMaker(operation string, columns []string, table string, conditions [][]string, values [][]string) (string, []interface{}) {
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
			for _, condition := range conditions {
				if len(condition) != 2 {
					return "invalid condition pair", nil
				}
				columnName := condition[0]
				value := condition[1]

				if !first {
					query += " AND "
				}
				// even if you use ? here, this query building function could have some SQL injection
				// you could pass `1 = 1; --` in columnName, and you select/delete everything
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
			for _, condition := range conditions {
				if len(condition) != 2 {
					return "invalid condition pair", nil
				}
				columnName := condition[0]
				value := condition[1]

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
		for i := 0; i < len(values); i++ {
			if len(values[i]) != 2 {
				return "invalid value pair", nil
			}
			if i < len(values)-1 {
				query += "?, "
			} else {
				query += "? "
			}
			args = append(args, values[i][1]) // Add the value part to args
		}
		query += ")"

		return query, args

	} else if operation == "delete" {
		query = fmt.Sprintf("DELETE FROM %s", table)

		if len(conditions) != 0 {
			query += " WHERE "
			first := true
			for _, condition := range conditions {
				if len(condition) != 2 {
					return "invalid condition pair", nil
				}

				columnName := condition[0]
				value := condition[1]

				if !first {
					query += " AND "
				}
				query += fmt.Sprintf("%s = ?", columnName)
				args = append(args, value)
				first = false
			}
		} else {
			return "no conditions provided for delete", nil
		}

		return query, args
	}

	return "invalid operation", nil
}

func InitializeAdminUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter password for your admin user")
	adminpassword, err := reader.ReadString('\n')
	// maybe the password has some space needed. I know I use some space in my passwords.
	adminpassword = strings.TrimSpace(adminpassword)
	if err != nil {
		fmt.Println(err)
	}
	//checking if the password is valid
	if ValidateUserInfoFormInputs("password", adminpassword) {
		//generating a user_id
		userId := tools.GenerateUUID()
		//hashing the password
		adminpassword = tools.HashThis(adminpassword)
		//creating a user record in users database
		// this querymaker API is not really readable
		query, arguments := QueryMaker("insert", []string{"userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{}, [][]string{{"userId", userId}, {"username", "admin"}, {"password", adminpassword}, {"firstName", "empty"}, {"lastName", "empty"}, {"email", "empty"}, {"phoneNumber", "empty"}, {"rule", "admin"}, {"suspended", "no"}})
		InitializeAdminUserInDatabase(query, arguments)
	} else {
		fmt.Println("Invalid password")
		os.Exit(0)
	}
}

func InitializeAdminUserInDatabase(query string, arguments []interface{}) {
	// not so safe as explained in the QueryMaker func
	safequery, err := DataBase.Prepare(query)
	if err != nil {
		// more than print, you should just return the error and stop the program, there is no good way forward after this.
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		// more than print, you should just return the error and stop the program, there is no good way forward after this.
		// and in the caller, you should check the error and stop there.
		fmt.Println(err)
	}
}

// use different function for different input, don't regroup everything. It's harder to read and reason about.
func ValidateUserInfoFormInputs(tobevalidated string, valuetobevalidated string) bool {
	validationFlag := true

	if tobevalidated == "username" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 || len(valuetobevalidated) > 30 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "password" {
		if tobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`[a-zA-Z]`).MatchString(valuetobevalidated) && regexp.MustCompile(`\d`).MatchString(valuetobevalidated) && regexp.MustCompile(`[\W_]`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "firstname" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile((`^[A-Za-z]+$`)).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "lastname" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[A-Za-z]+$`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "email" {
		if valuetobevalidated == "" {
			fmt.Println("Empty email")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 40 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "phonenumber" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) != 10 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`(^\d+$)`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "id" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`(^\d+$)`).MatchString(valuetobevalidated)) {
				validationFlag = false
				return validationFlag
			}
		}
	}

	return validationFlag
}

func ValidateTaskOrMessageInfoFormInputs(tobevalidated string, valuetobevalidated string) bool {
	validationFlag := true
	if tobevalidated == "priority" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 6 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile((`^[A-Za-z]+$`)).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "category" {
		if tobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			}
		}
	} else if tobevalidated == "title" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 30 {
				validationFlag = false
				return validationFlag
			}
		}
	} else if tobevalidated == "description" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 70 {
				validationFlag = false
				return validationFlag
			}
		}
	}
	return validationFlag
}

func ReadUser(query string, arguments []interface{}) []models.User {
	safequery, err := DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	user := models.User{}
	users := []models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Rule, &user.Suspended)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users
}
