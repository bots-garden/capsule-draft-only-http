package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/reactr/rwasm"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
### Start Capsule

```bash
CAPSULE_SETTINGS='{"token":"🙂🐼🍋🐯"}' \
CAPSULE_HTTP_PORT="8080" \
./capsule hello/hello.wasm
````

### Call the function

```bash
curl -X POST http://localhost:8080 \
	-H 'Content-Type: application/json' \
	-d '{"message":"Hello World", "author":"@k33g"}'; \
echo ""
````
*/

func main() {
	r := rt.New()

	// TODO: be able to download a Runnable
	mainWasmFunction := r.Register("main-function", rwasm.NewRunner(os.Args[1]))

	service := func(response http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)

		defer request.Body.Close()

		//{"parameters":{"message":"Hello World", "author":"@k33g"}, "settings":{"token":"🙂🐼🍋🐯"}}

		args := []byte(`{"parameters":` + string(body) + `,"settings":` + getEnv("CAPSULE_SETTINGS", "{}") +`}`)

		result, err := mainWasmFunction(args).Then()
		if err != nil {
			log.Println(err)
			return
		}

		response.WriteHeader(http.StatusOK)
		response.Write(result.([]byte))
	}

	listening_port := getEnv("CAPSULE_HTTP_PORT", "8080")

	http.HandleFunc("/", service)
	log.Println("🌍 Capsule is Listening on " + listening_port)
	http.ListenAndServe(":" + listening_port, nil)

}
