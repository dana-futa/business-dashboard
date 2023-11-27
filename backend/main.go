package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

type Employee struct {
	EmployeeID int    `json:"employeeId"` // EmployeeID starts with uppercase to make it an exported field
	FirstName  string `json:"firstName"`  // In json format, use proper camel case and convert between the 2 field formats
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Title      string `json:"title"`
	ManagerID  int    `json:"managerId"` // can be nil, default to -1
	IsActive   bool   `json:"isActive"`
}

var testEmployees = []Employee{
	{EmployeeID: 0, FirstName: "Bob", LastName: "Stevens", Email: "bob.stevens@business.com", Title: "CEO", IsActive: true}, // ManagerId: nil (Go will set this as default)
	{EmployeeID: 1, FirstName: "Mary", LastName: "Jane", Email: "mary.jane@business.com", Title: "Engineer 1", ManagerID: 0, IsActive: true},
	{EmployeeID: 2, FirstName: "John", LastName: "Hope", Email: "john.hope@business.com", Title: "Engineer 1", ManagerID: 0, IsActive: true},
	{EmployeeID: 3, FirstName: "David", LastName: "Gomez", Email: "david.gomez@business.com", Title: "Engineer 2", ManagerID: 1, IsActive: true},
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDatabaseConnection() error {
	db, err := sql.Open("sqlite", "../database/business.db")
	if err != nil {
		fmt.Println("An error in the db connection has occured.")
		return err
	}
	DB = db

	fmt.Println("Database connection successful!")

	return nil
}

/**
* Practice function to return testEmployees as a raw struct (dynamic array) to get the API working.
* @param context - information about the request
* @return testEmployees - slice
 */
func GetTestEmployees(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"data": testEmployees})
}

func GetAllActiveEmployees(context *gin.Context) {
	employees, err := GetAllActiveEmployeesHelper()
	checkErr(err)

	if employees == nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No Records Found."})
		return
	} else {
		context.IndentedJSON(http.StatusOK, gin.H{"data": employees})
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
		var managerID sql.NullInt64

		// bind the values in the row to the emp object
		err = rows.Scan(&emp.EmployeeID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.Title, &managerID, &emp.IsActive)

		// if there was an error binding the row data to the emp object, return nil for Employee struct and the error
		if err != nil {
			return nil, err
		}

		// set managerId on emp struct
		if managerID.Valid {
			emp.ManagerID = int(managerID.Int64)
		} else {
			emp.ManagerID = -1 // default of -1 to indicate no manager
		}

		// add extracted/mapped emp to employees slice
		employees = append(employees, emp)
	}

	return employees, err
}

func main() {
	fmt.Println("Hello world")

	// connect to the database and check if there was an error when connecting
	err := CreateDatabaseConnection()
	checkErr(err)

	// handles different routes/endpoints of the api
	router := gin.Default()

	// define API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test-employees", GetTestEmployees)
		v1.GET("/all-active-employees", GetAllActiveEmployees)
	}

	router.Run("localhost:8080")
}
