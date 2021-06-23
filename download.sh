#!/bin/bash

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' *.go

curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/board.map
curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/clarity.js
curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/index.html
curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/style.css
curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/server
curl -O https://raw.githubusercontent.com/xavi-/create-and-learn-6/main/run.sh

chmod +x server
chmod +x run.sh