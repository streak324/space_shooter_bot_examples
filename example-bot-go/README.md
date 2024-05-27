## Dependencies
- [TinyGo](https://github.com/tinygo-org/tinygo). build manually from dev branch
  - change `-bulk-memory` to `+bulk-memory` in line `features` in file: `targets/wasm-unknown.json`
  - follow instructions [here](https://tinygo.org/docs/guides/build/) for building
- [Binaryen](https://github.com/WebAssembly/binaryen) version 117 or higher
- Optional: [wabt](https://github.com/WebAssembly/wabt). you may have to remove lines in the build script that uses `wasm2wat` if you forgo it. 
- include both in your `PATH` env