package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type employee struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Age        string      `json:"age"`
	Employeers *employeers `json:"employeers"`
}

type employeers struct {
	Companyname string `json:"companyname"`
}

var Employee []employee

func getcompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Employee)
	
	


}

func getcompanybyid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range Employee {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}

}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Employee)
	params := mux.Vars(r)
	for index, item := range Employee {
		if item.Id == params["id"] {
			Employee = append(Employee[:index], Employee[index+1:]...)
			fmt.Println("deleted successfully")
			break
		}
	}
	json.NewEncoder(w).Encode(Employee)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	// emp.Id = strconv.Itoa(rand.Intn(1000000000))
	Employee = append(Employee, emp)
	json.NewEncoder(w).Encode(Employee)
}
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Employee {
		if item.Id == params["id"] {
			Employee = append(Employee[:index], Employee[index+1:]...)
			var emp employee
			_ = json.NewDecoder(r.Body).Decode(&emp)
			emp.Id = strconv.Itoa(rand.Intn(1000000000))
			Employee = append(Employee, emp)
			json.NewEncoder(w).Encode(Employee)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()
	// Employee = append(Employee, employee{Id: "1", Name: "mahesh", Age: "25", Employeers: &employeers{Companyname: "srm"}})
	// Employee = append(Employee, employee{Id: "2", Name: "mahe", Age: "25", Employeers: &employeers{Companyname: "srm"}})
	// Employee = append(Employee, employee{Id: "3", Name: "mahi", Age: "25", Employeers: &employeers{Companyname: "srm"}})
	r.HandleFunc("/get", getcompany).Methods("get")
	r.HandleFunc("/get/{id}", getcompanybyid).Methods("get")
	r.HandleFunc("/update/{id}", update).Methods("put")
	r.HandleFunc("/delete/{id}", delete).Methods("delete")
	r.HandleFunc("/create", create).Methods("post")
	log.Fatal(http.ListenAndServe(":8099", r))
}
