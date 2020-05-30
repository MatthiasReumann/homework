package main

func main(){
	srv := NewServer("8080")
	srv.Serve()
}
