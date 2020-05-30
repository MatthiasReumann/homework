package main

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type server struct{
	srv *http.Server
}

type HELink struct{
	HELinkUuid string
}

type HE struct{
	HELinkUuid string
	HeUuid string
}



func NewServer(port string) server{
	srv := &http.Server{
		Addr:           ":"+port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/links", Links)
	http.HandleFunc("/homeworks", Homeworks)

	return server{srv}
}

func Homeworks(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			var helink HELink

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error reading body: %v", err)
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(body, &helink)
			if err != nil{
				log.Print(err)
				http.Error(w, "Can't unmarshal data", http.StatusBadRequest)
			}

			uuid, err := uuid2.NewUUID()
			if err != nil{
				log.Fatal(err)
			}

			//TODO: check if helink is in db

			data := HE{helink.HELinkUuid, uuid.String()}

			//TODO: add HE to database

			json, err := json.Marshal(data)

			w.Write(json)
		default:
			w.Write([]byte("orsch"))
	}
}

func Links(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case http.MethodPost:
			uuid, err := uuid2.NewUUID()
			if err != nil{
				log.Fatal(err)
			}

			data := HELink{uuid.String()}

			//TODO:: Add helink to database

			json, err := json.Marshal(data)

			w.Write(json)
	default:
			w.Write([]byte("orsch"))
	}
}

func (s *server) Serve(){
	log.Printf("Listen at %s\n", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
