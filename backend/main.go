package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"

	"example.com/business-dashboard/constants"
	"example.com/business-dashboard/models"
)

func main() {
	fmt.Println("Hello world")

	// connect to the database and check if there was an error when connecting
	err := models.CreateDatabaseConnection()
	checkErr(err)

	// handles different routes/endpoints of the api
	router := gin.Default()

	// define API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test-employees", GetTestEmployees)
		v1.GET("/all-active-employees", GetAllActiveEmployees)
	}

	// set up CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{constants.FrontendURL}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// wrap router with CORS handler
	http.ListenAndServe(":8080", corsHandler(router))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* QUESTION: Although this is a small project, where would these functions go in an enterprise application?
Should they go in the employees.go model or is there typically a "Controller" and "Service" file like
in Java applications?

https://github.com/golang-standards/project-layout --> standard go project layout
From the authors: "If you are trying to learn Go or if you are building a PoC or a simple project for
yourself this project layout is an overkill. Start with something really simple instead
(a single main.gofile andgo.mod is more than enough)."
*/

/**
* Practice function to get the API working, which returns testEmployees as a raw struct (dynamic array).
* @param context - information about the request
* @return testEmployees - slice
 */
func GetTestEmployees(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"data": models.TestEmployees})
}

/**
* Practice function to get the API and DB connection returning, which returns all active employees in the database.
* @param context - information about the request
* @return if success --> {data: []Employee{}}
* 		  if error   --> {error: "error message"}
 */
func GetAllActiveEmployees(context *gin.Context) {
	employees, err := models.GetAllActiveEmployeesHelper()
	checkErr(err)

	if employees == nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No Records Found."})
		return
	} else {
		context.IndentedJSON(http.StatusOK, gin.H{"data": employees})
	}
}
