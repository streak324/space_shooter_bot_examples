@echo off
setlocal

:: Optional: Change the output file name
set "OUTPUT_NAME=bot"

:: do not use opt=z which is the default. may cause some unused wasm imports to be removed, which the wasm host would not like
tinygo build -o %OUTPUT_NAME%.wasm --opt 1 -target wasm-unknown .

rem Create destination directory if it doesn't exist
set DEST_DIR=%USERPROFILE%\AppData\Local\space_shooter_bots\scripts
if not exist "%DEST_DIR%" (
    mkdir "%DEST_DIR%"
)

rem Copy bot.wasm to the destination as go-example.wasm
copy /Y bot.wasm "%DEST_DIR%\go-bot.wasm"