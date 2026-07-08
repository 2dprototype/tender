@echo off
mkdir build
set GOARCH=amd64
set GOOS=windows
echo Building tender.exe (with OpenGL)...
go build -tags="gl glu glut glfw" -o build/tender.exe -ldflags "-s -w" cli/tender/main.go
echo Building tender_nogl.exe (without OpenGL)...
go build -o build/tender_nogl.exe -ldflags "-s -w" cli/tender/main.go
pause