#!/bin/bash -e

OUTPUT_NAME=bot

#do not use opt=z which is the default. may cause some unused wasm imports to be removed, which the wasm host would not like
tinygo build -o ${OUTPUT_NAME}.wasm --opt 1 -target wasm-unknown .
wasm2wat ${OUTPUT_NAME}.wasm -o ${OUTPUT_NAME}.wat