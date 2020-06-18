package main

func main() {
	db, err := NewDatabaseConnection("localhost", 5432, "homeexercise", "homeexercise", "he")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	srv := NewServer("3333", &db)
	srv.Serve()

}
