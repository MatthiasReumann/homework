package main

import "log"

func main() {
	db, err := NewDatabaseConnection("localhost", 5432, "homeexercise", "homeexercise", "he")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("connected to db!")
	}

	defer db.Close()

	srv := NewServer("3333", &db)
	srv.Serve()


}
