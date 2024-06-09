#!/bin/bash -e
#idk how to get clang to compile with imported wasm host funcs
clang --target=wasm32-unknown-unknown --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o bot.wasm main.c

#emcc main.c -o bot.wasm -s WASM=1 -s SIDE_MODULE=1
#emcc main.cpp -I./ -o bot.wasm -s WASM=1 -s INITIAL_HEAP=64kb -s STACK_SIZE=64kb -s MAXIMUM_MEMORY=512kb -s ALLOW_MEMORY_GROWTH=1 \
#	-Wl,--stack-first -Wl,--no-demangle -Wl,--no-entry -Wl,--import-memory

wasm2wat bot.wasm > bot.wat
wasm-objdump bot.wasm -d > bot.objdump



