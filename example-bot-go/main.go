package main

//export run
func run(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
