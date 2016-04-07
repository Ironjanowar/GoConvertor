#! /bin/bash

while true;do
    git pull
    go run bot.go
    sleep 1
done
