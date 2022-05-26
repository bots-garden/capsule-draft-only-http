package main

import (
	//"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/reactr/rwasm"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
// getEnv("CAPSULE_HTTP_PORT", "8080")

func main() {
	r := rt.New()

	mainWasmFunction := r.Register("main-function", rwasm.NewRunner(os.Args[1]))

	// use gin
	service := func (response http.ResponseWriter, request *http.Request)  {
		body, _ := ioutil.ReadAll(request.Body)

		defer request.Body.Close()

		// here: transform body to json, add environment variables
		// then json []byte

		result, err := mainWasmFunction(body).Then()
		if err != nil {
			fmt.Println(err) // log
			return
		}

		response.WriteHeader(http.StatusOK)
		response.Write(result.([]byte))

	}

	http.HandleFunc("/", service)
	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)

	// curl -X POST -d '{"FirstName": "Bob", "LastName": "Morane"}' http://localhost:8080


}
