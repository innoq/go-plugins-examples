#!/bin/bash

go build -o ./plugin/hello ./plugin/hello_impl.go
chmod +x ./plugin/hello
go build -o host .

echo "run me with ./host"