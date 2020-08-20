package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	PORTNUM = "8080"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./nc.jpg")
	io.WriteString(w, "Hello, NCSOFT!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("Started the golang-docker test server on port", PORTNUM, "Time :", time.Now())
	log.Fatal(http.ListenAndServe(":"+PORTNUM, nil))
}
