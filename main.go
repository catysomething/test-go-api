package main

import (
	"fmt"
	"net/http"
	"os"
)

// Go recognized functions that start with lowercase letter as
// private functions, only available within this pkg.

// by convention, variable name for writer is 'w' and
// variable name for reader is 'r'.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

func main() {
	//create default handler calling the function "rootHandler"
	http.HandleFunc("/", rootHandler)

	//spin up webserver and handle server any error
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
