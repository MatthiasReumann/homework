package main

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type homework_request struct {
	HELinkUuid string
	Firstname  string
	Lastname   string
}

type homeworkuuid_request struct {
	Text string
}

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func Links(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodPost:
		uuid, err := uuid2.NewUUID()
		if err != nil {
			log.Fatal(err)
		}

		data := HELink{uuid.String()}

		//TODO:: Add helink to database

		res, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error while marshalling: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(res)
	case http.MethodOptions:
		w.WriteHeader(200);
	default:
		MethodNotAllowed(w)
	}
}

func Homeworks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var request homework_request

		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Printf("Error Unmarshal body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//TODO: Check if helink-uuid is in database

		heuuid, err := uuid2.NewUUID()
		if err != nil {
			log.Fatal(err)
		}

		data := HE{
			request.HELinkUuid,
			heuuid.String(),
			Student{
				request.Firstname,
				request.Lastname},
			File{
				"", //TODO: Move to own constants file
			},
			HEStatusUnsubmitted}

		//TODO: Add HE to database

		res, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error while marshalling: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(res)
	case http.MethodOptions:
		w.WriteHeader(200);
	default:
		MethodNotAllowed(w)
	}
}

func HomeworksUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		heuuid := vars["uuid"]
		log.Print(heuuid)

		//TODO: check if heuuid is in db

		//TODO: load text from dataabase

		file := File{
			"text text",
		}

		res, err := json.Marshal(file)
		if err != nil {
			log.Printf("Error while marshalling: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(res)
	case http.MethodPut:
		vars := mux.Vars(r)
		heuuid := vars["uuid"]

		//TODO: check if heuuid is in db

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var request homeworkuuid_request

		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Printf("Error Unmarshal body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//TODO: update text via heuuid
		log.Printf("updated %s", heuuid)

		w.WriteHeader(http.StatusOK)
	default:
		MethodNotAllowed(w)
	}
}
