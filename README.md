# Capsule

### Start Capsule

```bash
CAPSULE_SETTINGS='{"token":"ILovePandas🐼"}' \
CAPSULE_HTTP_PORT="8080" \
./capsule hello/hello.wasm
```

### Call the Runnable's function

```bash
curl -X POST http://localhost:8080 \
	-H 'Content-Type: application/json' \
	-d '{"message":"Hello World", "author":"@k33g"}'; \
echo ""
```
