#!/bin/bash -e
npx asc assembly/index.ts --importMemory --use abort=assembly/index/myAbort -o build/bot.wasm
wasm2wat build/bot.wasm > build/bot.wat
