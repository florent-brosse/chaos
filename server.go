package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/scenarios", GetScenarios).Methods("GET")

	r.HandleFunc("/scenarios/{id}", GetScenario).Methods("GET")
	r.HandleFunc("/scenarios/{id}", CreateScenario).Methods("POST")
	//r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	listen := conf.Listen_address + ":" + strconv.Itoa(conf.Port)
	log.Fatal(http.ListenAndServe(listen, r))
}

func GetScenario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, v := range scenarios {
		if v.Id == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("scenario with id:'" + id + "' not found"))

}
func GetScenarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(scenarios)
}
func CreateScenario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var scenario Scenario
	_ = json.NewDecoder(r.Body).Decode(&scenario)
	scenario.Id = id
	for _, task := range scenario.Tasks {
		initDone(&task)
	}
	scenarios = append(scenarios, scenario)
	json.NewEncoder(w).Encode(scenario)
}

func initDone(task *Task) {
	task.Done = true
	task.Launched = true
	for _, tag := range task.Tags {
		if contains(conf.Tags, tag) {
			task.Done = false
			task.Launched = false
			return
		}
	}
	fmt.Println(task)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
