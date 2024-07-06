#!/bin/bash -e
#FLATCC_NO_ASSERT=1 or FLATCC_ASSERT=my_assert is required to avoid WASI imports
emcc main.c -DFLATCC_ASSERT=my_assert -I./ -o bot.wasm \
	-s WASM=1 -s INITIAL_HEAP=256kb -s STACK_SIZE=64kb \
	-s ALLOW_MEMORY_GROWTH=1 -s MAXIMUM_MEMORY=0 \
    -Wl,--stack-first -Wl,--no-demangle -Wl,--no-entry -Wl,--import-memory


wasm2wat bot.wasm > bot.wat
wasm-objdump bot.wasm -d > bot.objdump



