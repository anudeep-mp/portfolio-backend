package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anudeep-mp/portfolio-backend/router"
)

func main() {
	port := os.Getenv("PORT")
	r := router.Router()

	fmt.Println("Server is getting ready")
	http.ListenAndServe(":"+port, r)
	fmt.Printf("Listening at port %v", port)

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
