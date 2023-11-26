package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

type Employee struct {
	EmployeeId int    `json:"employeeId"` // EmployeeId starts with uppercase to make it an exported field
	FirstName  string `json:"firstName"`  // In json format, use proper camel case and convert between the 2 field formats
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Title      string `json:"title"`
	ManagerId  int    `json:"managerId"`
	IsActive   bool   `json:"isActive"`
}

var tempEmployees = []Employee{
	{EmployeeId: 0, FirstName: "Bob", LastName: "Stevens", Email: "bob.stevens@business.com", Title: "CEO", IsActive: true}, // ManagerId: nil (Go will set this as default)
	{EmployeeId: 1, FirstName: "Mary", LastName: "Jane", Email: "mary.jane@business.com", Title: "Engineer 1", ManagerId: 0, IsActive: true},
	{EmployeeId: 2, FirstName: "John", LastName: "Hope", Email: "john.hope@business.com", Title: "Engineer 1", ManagerId: 0, IsActive: true},
	{EmployeeId: 2, FirstName: "David", LastName: "Gomez", Email: "david.gomez@business.com", Title: "Engineer 2", ManagerId: 1, IsActive: true},
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func CreateDatabaseConnection() error {
// 	db, err := sql.Open("sqlite3", "business.db")
// 	if err != nil {
// 		return err
// 	}

// 	DB = db

// 	fmt.Println("Database connection successful!")

// 	defer DB.Close()

// 	return nil
// }

/**
* Practice function to return tempEmployees as a raw struct (dynamic array) to get the API working.
* @param context - information about the request
* @return tempEmployees - slice
 */
func GetTempEmployees(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tempEmployees)
}

func GetAllActiveEmployees(context *gin.Context) {
	employees, err := GetAllActiveEmployeesHelper()
	checkErr(err)

	if employees == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"data": employees})
	}
}

/**
* Practice function to get all employees in database to make sure db connection and API are working.
 */
func GetAllActiveEmployeesHelper() ([]Employee, error) {
	// query the database for all active Employees
	rows, err := DB.Query("SELECT * FROM Employees WHERE is_active IS TRUE")

	// if there was an error querying the database, return nil for Employee struct and the error
	if err != nil {
		return nil, err
	}

	// close database connection
	defer rows.Close()

	// create empty slice for employees
	employees := make([]Employee, 0)

	// iterate through rows of data retrieved from the database
	for rows.Next() {
		emp := Employee{}

		// bind the values in the row to the emp object
		err = rows.Scan(&emp.EmployeeId, &emp.FirstName, &emp.LastName, &emp.Email, &emp.Title, &emp.ManagerId, &emp.IsActive)

		// if there was an error binding the row data to the emp object, return nil for Employee struct and the error
		if err != nil {
			return nil, err
		}

		// add extracted/mapped emp to employees struct
		employees = append(employees, emp)
	}

	return employees, err
}

func main() {
	fmt.Println("Hello world")

	db, err := sql.Open("sqlite", "business.db")
	if err != nil {
		fmt.Println("An error in the db connection has occured")
		return
	}
	DB = db

	// handles different routes/endpoints of the api
	router := gin.Default()

	// // connect to the database and check if there was an error when connecting
	// err := CreateDatabaseConnection()
	// checkErr(err)

	// define API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/temp-employees", GetTempEmployees)
		v1.GET("/all-active-employees", GetAllActiveEmployees)
	}

	router.Run("localhost:8080")
}
