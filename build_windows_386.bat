@echo off
mkdir build
set GOARCH=386
set GOOS=windows
echo Building tender32.exe (with OpenGL)...
go build -tags="gl glu glut" -o build/tender32.exe -ldflags "-s -w" cli/tender/main.go
echo Building tender_nogl32.exe (without OpenGL)...
go build -o build/tender_nogl32.exe -ldflags "-s -w" cli/tender/main.go
pause