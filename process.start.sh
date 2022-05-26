#!/bin/bash
http_port=$1
wasm_module=$2

cmd="SAT_HTTP_PORT=${http_port} exec -a ${wasm_module}:${http_port} sat ${wasm_module}/${wasm_module}.wasm &"
# echo $cmd
read -p "$cmd"
eval "$cmd"
