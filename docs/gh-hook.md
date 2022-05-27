

```bash
subo create runnable gh-hook --lang tinygo

# create a github webhook with this url:
$(gp url 3000)
https://3000-botsgarden-capsule-9jlurr5ke1s.ws-eu44xl.gitpod.io
# Choose these options:
# - content type: application/json
# - Let me select individual events: issues + issue comments
# - Active == true


# run the webhook:
CAPSULE_SETTINGS='{"hook":"${SLACK_WASM_INCOMING_TOKEN}"}' \
CAPSULE_HTTP_PORT="3000" \
./capsule gh-hook/gh-hook.wasm

# Remark: don't forget to expose the 3000 http port (on Gitpod if you use Gitpod)
```

### Slack

- https://app.slack.com/client/T03HLS5KS49/C03H87T55CJ
- Add bot: https://slack.com/help/articles/115005265703-Create-a-bot-for-your-workspace

Settings > Administration > Manage Apps
Build
Create App
Incoming WebHook
Activate Incoming Webhooks

curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' ${SLACK_WASM_INCOMING_TOKEN}
