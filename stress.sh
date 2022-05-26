#!/bin/bash

concurrency=500
duration=30s
nbrequests=1000

# Invoke the hey load testing tool 
# with a concurrency of {concurrency} requests and 
# for a duration of {duration} seconds
hey -n ${nbrequests} -z ${duration} -c ${concurrency} -m POST -T "application/json" -d '{"message":"Hello World", "author":"@k33g"}' http://localhost:8080
