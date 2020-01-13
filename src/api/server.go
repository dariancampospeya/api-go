package api

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("Hola soy un comnetario nuevo")
	fmt.Println("\n\tServer listening[::]:3000")
	r := router.New()
	fmt.Println("\n\tServer listening[::]:3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
