package main

import (
	"github.com/suborbital/reactr/api/tinygo/runnable"
	"github.com/suborbital/reactr/api/tinygo/runnable/log"
	"github.com/tidwall/gjson"
)

type GhHook struct{}

func (h GhHook) Run(args []byte) ([]byte, error) {
	/* 🖐️ the structure of Json payload in `args`:
	   ```json
		 {
			 parameters: `the content of the request body (GitHub issue payload)`,
			 settings: `the settings of the Capsule isnstance (eg: token, ...)`
		 }
		 ```
	*/
	//log.Info("👋 " + string(args))

	// Get information about the issue
	//action := gjson.GetBytes(args, "parameters.issue.action")
	// eg: https://api.github.com/repos/bots-garden/sandbox/issues/1
	//url := gjson.GetBytes(args, "parameters.issue.url")
	//id := gjson.GetBytes(args, "parameters.issue.id")
	//number := gjson.GetBytes(args, "parameters.issue.number")
	title := gjson.GetBytes(args, "parameters.issue.title")
	//body := gjson.GetBytes(args, "parameters.issue.body")
	user := gjson.GetBytes(args, "parameters.issue.user.login")

	log.Info("📝: " + title.Str + " by " + user.Str)

	return []byte("Hello 🐙"), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(GhHook{})
}
