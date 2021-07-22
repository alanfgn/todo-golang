package main

import (
	"log"
	"net/http"
	"os"
//	"eddy.com/todo/data"
	"eddy.com/todo/route"
)

func main() {
	var PORT = ":9000"
	//data.InitDatabase()

	if envPort := os.Getenv("PORT"); envPort != "" {
		PORT = ":" + envPort
	}

	log.Print("Server on port: ", PORT)
	http.ListenAndServe(PORT, route.RegisterRoute())
}
