package main

import (

	"github.com/suborbital/reactr/api/tinygo/runnable"
	"github.com/suborbital/reactr/api/tinygo/runnable/http"
	"github.com/suborbital/reactr/api/tinygo/runnable/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

)

type GhHook struct{}

// Triggered by a GitHub WebHook
func (h GhHook) Run(args []byte) ([]byte, error) {
	/* 🖐️ the structure of Json payload in `args`:
	   ```json
		 {
			 parameters: `the content of the request body (GitHub issue payload)`,
			 settings: `the settings of the Capsule isnstance (eg: token, ...)`
		 }
		 ```
	*/

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

	slackMessage, _ := sjson.Set(`{"text":""}`, "text", "📝: " + title.Str + " by " + user.Str)

	headers := make(map[string]string)
	headers["Content-type"] = "application/json; charset=utf-8"

	hookUrl := gjson.GetBytes(args, "settings.hook")

	log.Info("🌍: " + hookUrl.Str)
	log.Info("📝: " + slackMessage)
	log.Info("🤯: " + headers["Content-type"])
	
	// Post message to Slack
	http.POST(hookUrl.Str, []byte(slackMessage), headers)

	return []byte("Hello 🐙"), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(GhHook{})
}
