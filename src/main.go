package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pbentes/80_20/src/db"
	"github.com/pbentes/80_20/src/router"
)

func main() {
	godotenv.Load()

	db.InitDB()
	defer db.Cleanup()

	mux := router.SetupRouter()

	fmt.Println("Server running on: http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
