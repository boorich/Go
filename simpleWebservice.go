package main // so this can be executed from the cli

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // any request to this service is using this handler, return is a (anonymous) function
		names := r.URL.Query()["name"] // return map of data with all query params the app accepted and retun anything matching "name"
		var name string
		if len(names) == 1 { //since this might return more than one name
			name = names[0]
		}
		m := map[string]string{"name": name} //map object from string to string that can easily be serialized into a JSON string initialized to "name":name
		enc := json.NewEncoder(w)            // pass writer interface to the new encoder
		enc.Encode(m)                        // encode the map

	})

	err := http.ListenAndServe(":3000", nil) // nil makes go's standard handler run, http:ListenAndServe make err of type error

	if err != nil { // something bad happened
		log.Fatal(err)
	}

}
