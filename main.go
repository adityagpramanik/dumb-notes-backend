package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"dumbnotes.com/dumb-notes-backend/routes"
	"github.com/joho/godotenv"
)

// this runs before main, project initialisations can be added here
func init() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading environment secrets");
	}
}

func main()  {
	router := routes.SetupRoutes()
	PORT := os.Getenv("PORT")
	fmt.Printf("Server listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}