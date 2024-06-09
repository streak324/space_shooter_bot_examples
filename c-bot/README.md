# dependancies
 - [emscripten](https://emscripten.org/docs/getting_started/index.html)
# NOTES
 - make sure all your global variables and non-exported funcs are `static`. Otherwise they get included as wasm imports in the which will not be supported by the wasm host

