#!/bin/bash

go build -buildmode=plugin -o hello/plugin.so hello/greeter.go
go build -o host
chmod +x ./host

echo "run me with ./host"