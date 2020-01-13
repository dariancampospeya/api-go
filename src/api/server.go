package api

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("\n\tServer On..[::]:3000")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}
