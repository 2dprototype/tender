@echo off
mkdir build
set GOARCH=amd64
set GOOS=darwin
echo Building tender_darwin.exe...
go build -o build/tender_darwin.exe -ldflags "-s -w" cli/tender/main.go
pause