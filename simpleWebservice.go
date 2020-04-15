package main // so this can be executed from the cli

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // any request to this service is using this handler, return is a (anonymous) function
		w.Write([]byte("Hello World")) // make Hello World a byte array so you can the writer interface of the response writer
	})

	err := http.ListenAndServe(":3000", nil) // nil makes go's standard handler run, http:ListenAndServe make err of type error

	if err != nil { // something bad happened
		log.Fatal(err)
	}

}
