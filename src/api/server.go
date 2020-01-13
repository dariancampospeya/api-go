package api

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {

	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))

}
