

## Run Sat with a process name

```bash
SAT_HTTP_PORT=8080 exec -a demo-go sat demo-go/demo-go.wasm &
SAT_HTTP_PORT=8081 exec -a demo-rust sat demo-rust/demo-rust.wasm &
```



> Get the list of the sat processes
```bash
ps -fC sat

# output
gitpod      2501       1  0 11:11 ?        00:00:00 demo-go demo-go/demo-go.wasm
gitpod      2875    2612  0 11:11 pts/1    00:00:00 demo-rust demo-rust/demo-rust.wasm
```

> Kill a sat process by name
```bash
pkill -f demo-go
pkill -f demo-rust
```


