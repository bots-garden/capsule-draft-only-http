#!/bin/bash
process_name=$1

cmd="pkill -f ${process_name}"
# echo $cmd
read -p "$cmd"
eval "$cmd"
