#!/bin/bash -e
#idk how to get clang to compile with imported wasm host funcs
#clang --target=wasm32-unknown-unknown --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o bot.wasm main.c
emcc main.c -o bot.wasm -s WASM=1 -s SIDE_MODULE=1
wasm2wat bot.wasm > bot.wat



