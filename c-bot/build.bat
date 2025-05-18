@echo off
rem FLATCC_NO_ASSERT=1 or FLATCC_ASSERT=my_assert is required to avoid WASI imports
call emcc main.c -DFLATCC_ASSERT=my_assert -I./ -o bot.wasm ^
    -gsource-map ^
    -s WASM=1 -s INITIAL_HEAP=256kb -s STACK_SIZE=64kb -s ALLOW_MEMORY_GROWTH=1 -s MAXIMUM_MEMORY=0 ^
    -Wl,--stack-first -Wl,--no-demangle -Wl,--no-entry -Wl,--import-memory

rem Create destination directory if it doesn't exist
set DEST_DIR=%USERPROFILE%\AppData\Local\space_shooter_bots\scripts
if not exist "%DEST_DIR%" (
    mkdir "%DEST_DIR%"
)

rem Copy bot.wasm to the destination
copy /Y bot.wasm "%DEST_DIR%"
