## gh-rust-hook

- This Runnable is triggered by a GitHub WebHook when you create an issue. 
- When a GitHub issue is created, the Runnable posts a message to a Slack "Incoming WebHook"

### Slack Incoming WebHook

You need to create an incoming webhook:

- Settings > Administration > Manage Apps
- Build
- Create App > Incoming WebHook
- Activate Incoming Webhooks

Get the Incoming WebHook URL and set an environment variable with the value (eg: `SLACK_WASM_INCOMING_TOKEN`)

> Post a message to Slack:
```bash
curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' ${SLACK_WASM_INCOMING_TOKEN}
```

### GitHub WebHook

You must provide the url of the service to the GitHub WebHook.

> if you use Gitpod
```bash
# Get an URL for the service listening on the 3000 port
$(gp url 3000) 
# create a github webhook with the url provided by Gitpod
# eg:
# https://3000-botsgarden-capsule-qavlryglfdp.ws-eu44xl.gitpod.io
#
# Choose these options:
# - content type: application/json
# - select individual events: issues 
# - Active == true
#
# Remark: don't forget to expose the 3000 http port on Gitpod (see `.gitpod.yml`)
```

### Run the external service

```bash
CAPSULE_SETTINGS='{"hook":"'"${SLACK_WASM_INCOMING_TOKEN}"'"}' \
CAPSULE_HTTP_PORT="3000" \
./capsule gh-rust-hook/gh-rust-hook.wasm
```

> The main code of **capsule** is located here: ../main.go
