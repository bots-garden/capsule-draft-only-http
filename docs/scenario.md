# Scenario

## scketchnotes

- What is Wasm? + History
- What are Suborbital, Reactr, Sat, Subo (and more)

## install Sat

> show the Dockerfile

## 1st demo (local wasm module)

```bash
subo create runnable demo-go --lang tinygo
# explain and change the source code
cd demo-go
subo build .
cd ..
SAT_HTTP_PORT=8080 sat demo-go/demo-go.wasm
curl -d 'Bob Morane' http://localhost:8080; echo ""
```

## 2nd demo (download wasm module from https location)

```bash
node serve.wasm.js
# in another terminal
subo create runnable demo-swift --lang swift
# explain and change the source code
cd demo-swift
subo build .
cd ..
SAT_HTTP_PORT=8080 sat $(gp url 3000)/demo-swift/demo-swift.wasm
curl -d 'Bob Morane' http://localhost:8080; echo ""
```

## 3rd demo (host functions)

```bash
subo create runnable demo-rust
# explain and change the source code
# use the snippet
cd demo-rust
subo build .
cd ..
node hello.js
# in another terminal
SAT_HTTP_PORT=8080 sat demo-rust/demo-rust.wasm
curl -d 'Bob Morane' http://localhost:8080; echo ""
```

## 4th demo (JavaScript function)

> 🚧 work in progress