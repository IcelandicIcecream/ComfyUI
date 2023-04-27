#!/bin/bash -i
wget -q -O - https://git.io/vQhTU | bash -s -- --version 1.19
source ~/.bashrc
go run main.go
