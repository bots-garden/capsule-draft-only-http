package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/reactr/rwasm"
	"github.com/suborbital/vektor/vk"
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
CAPSULE_SETTINGS='{"token":"ILovePandas🐼"}' \
CAPSULE_HTTP_PORT="8080" \
./capsule hello/hello.wasm
```

### Call the function

```bash
curl -X POST http://localhost:8080 \
	-H 'Content-Type: application/json' \
	-d '{"message":"Hello World", "author":"@k33g"}'; \
echo ""
```
*/

func main() {
	r := rt.New()

	// TODO: be able to download a Runnable
	mainWasmFunction := r.Register("main-function", rwasm.NewRunner(os.Args[1]))

	listening_port, _ := strconv.Atoi(getEnv("CAPSULE_HTTP_PORT", "8080"))
	server := vk.New(
		vk.UseAppName("🌍 Capsule server"),
		vk.UseHTTPPort(listening_port),
	)

	// The handler is triggered at every POST request
	wasmHandler := func(request *http.Request, ctx *vk.Ctx) (interface{}, error) {
		body, _ := ioutil.ReadAll(request.Body)

		defer request.Body.Close()

		//{"parameters":{"message":"Hello World", "author":"@k33g"}, "settings":{"token":"ILovePandas🐼"}}

		/*
			1- Create a Json String with 2 fields:
			   - `parameters` (the value is the content of `request.Body`)
				 - `settings`` (the value is the content of the `CAPSULE_SETTINGS` environment variable)
			2- Convert the Json String to a `[]byte`
			  - the result will be the argument of the `Run` method of the Runnable
		*/
		args := []byte(`{"parameters":` + string(body) + `,"settings":` + getEnv("CAPSULE_SETTINGS", "{}") + `}`)

		// Call the Wasm function
		result, err := mainWasmFunction(args).Then()
		if err != nil {
			log.Println(err)
			return err, nil
		}
		return result, nil
	}

	server.POST("/", wasmHandler)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
