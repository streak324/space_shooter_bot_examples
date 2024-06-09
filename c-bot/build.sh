#!/bin/bash -e
#idk how to get clang to compile with wasm
#clang --target=wasm32-unknown-unknown --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o bot.wasm main.c

emcc main.c -I./ -o bot.wasm -s WASM=1 -s INITIAL_HEAP=256kb -s STACK_SIZE=64kb -s ALLOW_MEMORY_GROWTH=0 \
	-Wl,--stack-first -Wl,--no-demangle -Wl,--no-entry -Wl,--import-memory

wasm2wat bot.wasm > bot.wat
wasm-objdump bot.wasm -d > bot.objdump



