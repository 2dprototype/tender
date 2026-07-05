@echo off
mkdir build
set GOARCH=amd64
set GOOS=linux
echo Building tender_linux.exe...
go build -o build/tender_linux.exe -ldflags "-s -w" cli/tender/main.go
pause