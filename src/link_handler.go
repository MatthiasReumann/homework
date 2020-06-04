package main

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"log"
	"net/http"
)

type HELink struct {
	HELinkUuid string
}

func Links(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK);
	default:
		MethodNotAllowed(w)
	}
}
