// add.c
#include <stdint.h>

typedef int32_t i32;
typedef int8_t i8;

extern void botsLog(void* buffer_ptr, i32 length);

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
	for (i = 0; i < string_length; i++) {
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
void step()
{
	step_count += 1;
	reset_buffer(&print_buffer);
	write_string_to_buffer(&print_buffer, (const i8*) "Hello World. I am C\n\0", 0);
	botsLog(print_buffer.ptr, print_buffer.offset);
}

