## Dependencies
- [TinyGo](https://github.com/tinygo-org/tinygo) 0.31.2 or higher
- [Binaryen](https://github.com/WebAssembly/binaryen) version 117 or higher
- Optional: [wabt](https://github.com/WebAssembly/wabt). you may have to remove lines in the build script that uses `wasm2wat` if you forgo it. 
- include both in your `PATH` env