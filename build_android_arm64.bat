@echo off
mkdir build
set GOARCH=arm64
set GOOS=android
echo Building tender_termux.exe...
go build -o build/tender_termux.exe -ldflags "-s -w" cli/tender/main.go
pause