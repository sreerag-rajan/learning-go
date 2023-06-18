package main

import (
	"fmt"
	"log"
	"net/http"

	"github.coms/sreerag-rajan/mongoapi/router"
)

func main() {

	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started ....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on Port 4000")
}

// mongodb+srv://sreeragrajan5:mTGZxTKQo1Umkbo1@cluster0.1hhrcat.mongodb.net/?retryWrites=true&w=majority
