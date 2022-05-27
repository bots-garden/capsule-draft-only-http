

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
CAPSULE_SETTINGS='{"token":"ILovePandas🐼"}' \
CAPSULE_HTTP_PORT="3000" \
./capsule gh-hook/gh-hook.wasm

# Remark: don't forget to expose the 3000 http port (on Gitpod if you use Gitpod)
```
