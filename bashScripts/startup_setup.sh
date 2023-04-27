#!/bin/bash -i
wget -q -O - https://git.io/vQhTU | bash -s -- --version 1.19
source /root/.bashrc | bash
go run main.go
