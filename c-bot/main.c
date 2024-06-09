#include <stdint.h>
#include <emscripten.h>

typedef int32_t i32;
typedef int8_t i8;


//any imported wasm host functions need the function declared with EM_IMPORT(func_name) appended to it. no need to define function since it is provided by the wasm host
void botsLog(void* buffer_ptr, i32 length) EM_IMPORT(botsLog);

typedef struct {
	i8* ptr;
	i32 offset;
	i32 capacity;
} Buffer;

#define PRINT_BUFFER_CAPACITY 4096
static i8 _print_buffer[PRINT_BUFFER_CAPACITY];
static Buffer print_buffer = {
	_print_buffer,
	0,
	PRINT_BUFFER_CAPACITY
};

// string must be either null terminated or passed in a string length
static void write_string_to_buffer(Buffer* buffer, const i8* string_ptr, const i32 string_length)
{
	i32 i = 0;
	for (i = 0; string_length == 0 || i < string_length; i++) {
		if (buffer->offset >= buffer->capacity) {
			return;
		}

		i8 b = string_ptr[i];	
		if (b == 0) {
			return;
		};
		buffer->ptr[buffer->offset] = b;
		buffer->offset += 1;
	}
}

static void reset_buffer(Buffer* buffer) {
	buffer->offset = 0;
}

static i32 step_count = 0;

EMSCRIPTEN_KEEPALIVE void step()
{
	step_count += 1;
	reset_buffer(&print_buffer);
	write_string_to_buffer(&print_buffer, (const i8*) "Hello World. I am C\n\0", 0);
	botsLog(print_buffer.ptr, print_buffer.offset);
}

