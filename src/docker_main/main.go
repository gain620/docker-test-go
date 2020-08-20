package main

import (
	"io"
	"net/http"
	"log"
	"time"
	"fmt"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./nc.jpeg")
	io.WriteString(w, "Hello, NCSOFT!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("Started the golang-docker test server! Time :", time.Now())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
