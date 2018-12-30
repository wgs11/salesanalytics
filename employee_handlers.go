package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	ID 	int 	`json: "id"`
	FN	string	`json: "fn"`
	LN	string	`json: "ln"`
	Role	string	`json: "role"`
}



type EmployeesStruct struct{
	PageTitle string
	Person Employee

}


func stringToInt(id string) int {
	fval, err := strconv.ParseInt(id,0, 0)
	if err != nil {
		fmt.Println(err)
	}
	return int(fval)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(w,r) {
		fmt.Println("admin check failed")
		w.WriteHeader(http.StatusForbidden)
	} else {
		user := getUser(w,r)
		fmt.Println("got here")
		if (user == "") {
			log.Fatal("No user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err := templates.ExecuteTemplate(w, "admin", user)
		if err != nil {
			fmt.Println(fmt.Errorf("Error %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.FormValue("id")
	employees, err := store.GetEmployee(stringToInt(id))
	Employees := EmployeesStruct{}
	fmt.Println(len(employees))
	Employees.Person = (*employees[0])
	Employees.PageTitle = "This is a Title"
	err = templates.ExecuteTemplate(w,"employee",Employees)
	if err != nil {
		log.Fatal("Cannot retrive page.", err)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
