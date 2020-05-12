#!/bin/bash 
export GOOS=windows
export GOARCH=386
go build -o ./bin/web-server.exe main.go


export GOOS=linux
export GOARCH=amd64
go build -o ./bin/web-server main.go