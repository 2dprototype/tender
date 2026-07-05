@echo off
mkdir build

echo === Building Windows x64 ===
set GOARCH=amd64
set GOOS=windows
echo Building tender.exe (with OpenGL)...
go build -tags="gl glu glut" -o build/tender.exe -ldflags "-s -w" cli/tender/main.go
echo Building tender_nogl.exe (without OpenGL)...
go build -o build/tender_nogl.exe -ldflags "-s -w" cli/tender/main.go

echo === Building Windows x86 (32-bit) ===
set GOARCH=386
set GOOS=windows
echo Building tender32.exe (with OpenGL)...
go build -tags="gl glu glut" -o build/tender32.exe -ldflags "-s -w" cli/tender/main.go
echo Building tender_nogl32.exe (without OpenGL)...
go build -o build/tender_nogl32.exe -ldflags "-s -w" cli/tender/main.go

echo === Building Linux x64 ===
set GOARCH=amd64
set GOOS=linux
echo Building tender_linux.exe...
go build -o build/tender_linux.exe -ldflags "-s -w" cli/tender/main.go

echo === Building Android / Termux ARM64 ===
set GOARCH=arm64
set GOOS=android
echo Building tender_termux.exe...
go build -o build/tender_termux.exe -ldflags "-s -w" cli/tender/main.go

echo === Building Darwin x64 ===
set GOARCH=amd64
set GOOS=darwin
echo Building tender_darwin.exe...
go build -o build/tender_darwin.exe -ldflags "-s -w" cli/tender/main.go

echo All builds finished!
pause
