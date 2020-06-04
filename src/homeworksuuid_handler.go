package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HomeworksUUID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		homeworkuuid_get(w, r)
		return
	case http.MethodPut:
		homeworkuuid_put(w, r)
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK);
		return
	default:
		MethodNotAllowed(w)
	}
}

func homeworkuuid_get(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	heuuid := vars["uuid"]
	log.Print(heuuid)

	//TODO: check if heuuid is in db  //TODO: load text from dataabase

	file := File{
		"text text",
	}

	res, err := json.Marshal(file)
	if err != nil {
		log.Printf("Error while marshalling: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func homeworkuuid_put(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	heuuid := vars["uuid"]

	//TODO: check if heuuid is in db

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := struct {
		Text string
	}{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Error Unmarshal body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: update text via heuuid

	log.Printf("updated %s", heuuid)

	w.WriteHeader(http.StatusOK)
}