package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmployeeId   string `json:"employeeid"`
	EmployeeName string `json:"employeename"`
	Company      string `json:"company"`
	Salary       int    `json:"salary"`
}

var employees = []Employee{
	{
		EmployeeId:   "001",
		EmployeeName: "Heinner",
		Company:      "Truelogic",
		Salary:       100,
	},
	{
		EmployeeId:   "002",
		EmployeeName: "Braham",
		Company:      "Zaga",
		Salary:       200,
	},
}

// Fetch all Employes Data
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// Create Employee with Data
func createEmployee(c *gin.Context) {
	var newEmployee Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		meesageOutput := fmt.Sprintf("Error on user creation >> %v", err.Error())
		c.JSON(400, gin.H{
			"message": meesageOutput,
		})

		return
	}

	// check if user exist
	for _, user := range employees {
		if user.EmployeeName == newEmployee.EmployeeName || user.EmployeeId == newEmployee.EmployeeId {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("User Already exist >> %v", newEmployee),
			})

			return
		}
	}

	// user adding
	employees = append(employees, newEmployee)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("User Created >> %v", newEmployee),
	})
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/employees", getEmployees)
	r.POST("/employees", createEmployee)

	r.Run() // listen and serve on 0.0.0.0:8080
}
