#!/bin/bash -e
clang --target=wasm32 --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o c-example.wasm main.c


