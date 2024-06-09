#include <stdint.h>
#include "game_state_reader.h"
#include <emscripten.h>

typedef int8_t i8;
typedef int32_t i32;
typedef int64_t i64;

typedef uint64_t u64;

typedef float f32;
typedef double f64;

#undef ns
#define ns(x) FLATBUFFERS_WRAP_NAMESPACE(MyGameState, x) // Specified in the schema.

//any imported wasm host functions need the function declared with EM_IMPORT(func_name) appended to it. no need to define function since it is provided by the wasm game host
void botsLog(void* buffer_ptr, i32 length) EM_IMPORT(botsLog);
i32 botsGetGameState(void* buffer_ptr, i32 capacity) EM_IMPORT(botsGetGameState);
i32 botsMoveEntityToTarget(u64 entity_id, f32 x, f32 y) EM_IMPORT(botsMoveEntityToTarget);

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

static Buffer game_state_buffer = {};

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

typedef struct {
	f32 x;
	f32 y;
} Vec;

Vec start_position = {};

EMSCRIPTEN_KEEPALIVE void step()
{
	if (step_count == 0) {
		const i32 game_state_buffer_capacity = 8 * 1024;
		game_state_buffer.ptr = (i8*) malloc(game_state_buffer_capacity);
		game_state_buffer.offset = 0;
		game_state_buffer.capacity = game_state_buffer_capacity;
	}
	GameState_table_t game_state = GameState_as_root((void*) game_state_buffer.ptr);
	Entity_vec_t entities = GameState_entities_get(game_state);
	//size_t num_entities = Entity_vec_len(entities);

	//i32 i = 0;
	//for (i = 0; i < num_entities; i++) {
	//	Entity_table_t entity = Entity_vec_at(entities, i);
	//	flatbuffers_bool_t is_mine = Entity_my_get(entity);
	//	if (is_mine) {
	//		Vec2_table_t position = Entity_position_get(entity);
	//		if (step_count == 0) {
	//			start_position.x = Vec2_x_get(position);
	//			start_position.y = Vec2_y_get(position);
	//		}
	//		botsMoveEntityToTarget(Entity_id_get(entity), 0.0f, 0.0f);
	//	}
	//}

	reset_buffer(&print_buffer);
	write_string_to_buffer(&print_buffer, (const i8*) "Hello World. I am C++\n\0", 0);
	botsLog(print_buffer.ptr, print_buffer.offset);

	step_count += 1;
}

