package main

import (
	"fmt"
	"net/http"
	"runtime"
	"encoding/json"
	"io/ioutil"
)

const PORT = ":4567"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		var person map[string]interface{}

		if err := json.Unmarshal(body, &person); err != nil {
			panic(err)
		}

		fmt.Println(person["email"])
	}
}

func main() {
	http.HandleFunc("/person", httpHandler)
	error := http.ListenAndServe(PORT, nil)
	if error != nil {
		panic(error)
	}
}

func init() {
	// Use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
}
