#!/bin/bash -e
#idk how to get clang to compile with imported wasm host funcs
#clang --target=wasm32-unknown-unknown --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o bot.wasm main.c
emcc main.c -o bot.wasm -c
wasm2wat bot.wasm > bot.wat



