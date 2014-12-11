package main

import (
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
