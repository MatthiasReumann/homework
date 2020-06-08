package main

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Link struct {
	Uuid string
}

type List struct{
	Link Link
	Submissions []string
}

func (s *server) Links(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		uuid, err := uuid2.NewUUID()
		if err != nil {
			log.Fatal(err)
		}

		data := Link{uuid.String()}

		// add link to db
		err = env.db.AddHelink(data.Uuid)
		if err != nil {
			log.Printf("Error: could not connect to db")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error while marshalling: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(res)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK);
	default:
		s.MethodNotAllowed(w)
	}
}

func (s *server) LinksUUID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		uuid := vars["uuid"]
		log.Print(uuid)

		//TODO: Get all homework uuids via link

		helist := []string{"1","3","3"}

		data := List{Link{uuid}, helist}

		res, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error while marshalling: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(res)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK);
	default:
		s.MethodNotAllowed(w)
	}
}

