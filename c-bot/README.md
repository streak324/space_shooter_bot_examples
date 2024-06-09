# dependancies
- [emscripten](https://emscripten.org/docs/getting_started/index.html)
  - version `3.1.61` has been tested to work in the example, but other versions probably work too
# NOTES
- make sure all your global variables and non-exported funcs are `static`. Otherwise they get included as wasm imports in the which will not be supported by the wasm host

