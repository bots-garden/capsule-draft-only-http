#!/bin/bash
runnable_name=$1
language=$2

cmd="subo create runnable ${runnable_name} --lang ${language}"
# echo $cmd
read -p "$cmd"

eval "$cmd"
