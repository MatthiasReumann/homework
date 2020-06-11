package main

import "log"

type Env struct {
	db databaseConnection
}

var env *Env

func main() {
	db, err := NewDatabaseConnection("localhost", 5432, "homeexercise", "homeexercise", "he")

	env = &Env{db} // temporarily

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("connected to db!")
	}

	defer db.Close()

	srv := NewServer("3333")
	srv.Serve()

}
