@echo off
setlocal

:: Set the target OS and architecture
set "GOOS=js"
set "GOARCH=wasm"

:: Optional: Change the output file name
set "OUTPUT_NAME=bot"

:: Build the Go application
echo Building for %GOOS%/%GOARCH%
set GOOS=%GOOS%
set GOARCH=%GOARCH%
tinygo build -o %OUTPUT_NAME%.wasm -target wasm-unknown .
wasm2wat %OUTPUT_NAME%.wasm -o %OUTPUT_NAME%.wat