#!/bin/bash

payload=$(</dev/stdin)

echo "received payload:"
echo $payload | jq '.' | cat
