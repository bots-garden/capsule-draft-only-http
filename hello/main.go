package main

import (
	"github.com/suborbital/reactr/api/tinygo/runnable" 
	//"github.com/suborbital/reactr/api/tinygo/runnable/http" 
	"github.com/suborbital/reactr/api/tinygo/runnable/log" 
	"github.com/tidwall/gjson"
)

type Hello struct{}


func (h Hello) Run(args []byte) ([]byte, error) {

	//http.Get("")
	//log.Info("👋 " + string(args))

	token := gjson.GetBytes(args, "settings.token")
	
	log.Info("🔐: " + token.Str)

	message := gjson.GetBytes(args, "parameters.message")
	author := gjson.GetBytes(args, "parameters.author")

	log.Info("📧: " + message.Str)
	log.Info("👨‍🏫: " + author.Str)

	return []byte("👋 " + message.Str + " from " + author.Str), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(Hello{})
}
