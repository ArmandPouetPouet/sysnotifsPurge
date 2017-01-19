package main

import (
	routes "SysnotifsPurge/Routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":80", router))
}
