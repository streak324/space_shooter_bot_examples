## Dependencies
- [emscripten](https://emscripten.org/docs/getting_started/index.html)
  - version `3.1.61` has been tested to work in the example, but other versions probably work too
- Clang 14+
### Ubuntu
- `sudo apt install gcc-multilib`

## NOTES
- you may need all your global variables and non-exported funcs to be `static`. Otherwise they may get included as wasm imports which will not be supported by the wasm host
- Avoid using any libraries that requires access to operating system or accesses any syscalls or filesystem.

