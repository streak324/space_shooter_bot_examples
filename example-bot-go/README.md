## Dependencies
- [TinyGo](https://github.com/tinygo-org/tinygo). build manually from dev branch
  - change `-bulk-memory` to `+bulk-memory` in line `features` in file: `targets/wasm-unknown.json`
  - follow instructions [here](https://tinygo.org/docs/guides/build/) for building
  - if trying to build TinyGo on Windows, you will have an easier time using Windows Subsystem for Linux instead
- [Binaryen](https://github.com/WebAssembly/binaryen) version 117 or higher
- Optional: [wabt](https://github.com/WebAssembly/wabt). you may have to remove lines in the build script that uses `wasm2wat` if you forgo it. 
- include dependencies in your `PATH` env

## Troubleshooting
- avoid `fmt`, `os`, `reflect`, `encoding/*` and any other packages dependent on them. they are not supported on TinyGo `wasm-unknown` target.
- `encoding/json` won't work, so some other tool is needed for json parsing. possible minor modifications needed to make a third party json parsing library work with the `wasm-unknown` compile target
  - see gjson folder for example, which is borrowed from [here](https://github.com/tidwall/gjson)
- avoid strconv.FormatFloat. it can cause too many allocations