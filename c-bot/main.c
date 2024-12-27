#define my_assert(expr) if (!(expr)) { *(volatile int*)0 = 0; }

#include <stdint.h>
#include "game_state_reader.h"
#include <emscripten.h>
#include <stdio.h>
#include <math.h>
#include <stdarg.h>

typedef int8_t i8;
typedef int32_t i32;
typedef int64_t i64;

typedef uint8_t u8;
typedef uint32_t u32;
typedef uint64_t u64;

typedef float f32;
typedef double f64;

//any imported wasm host functions need the function declared with EM_IMPORT(func_name) appended to it. no need to define function since it is provided by the wasm game host
void botsLog(void* buffer_ptr, i32 length) EM_IMPORT(botsLog);
i32 botsGetGameState(void* buffer_ptr, i32 capacity) EM_IMPORT(botsGetGameState);
i32 botsMoveEntityToTarget(u64 entity_id, f32 x, f32 y) EM_IMPORT(botsMoveEntityToTarget);
i32 botsAimTurret(u64 entity_id, i32 block_index, f32 x, f32 y) EM_IMPORT(botsAimTurret);
i32 botsFireCannon(u64 entity_id, i32 block_index) EM_IMPORT(botsFireCannon);
u32 botsFindPath(f32 start_x, f32 start_y, f32 goal_x, f32 goal_y, void *buffer_ptr, u32 capacity) EM_IMPORT(botsFindPath);
i32 botsGrabFlag(u64 entityId) EM_IMPORT(botsGrabFlag);
void botsDrawText(void* buffer_ptr, u32 length, f32 x, f32 y, f32 size, u32 color) EM_IMPORT(botsDrawText);
void botsDrawLine(f32 x1, f32 y1, f32 x2, f32 y2, f32 width, u32 color) EM_IMPORT(botsDrawLine);
void botsDrawRectangle(f32 x, f32 y, f32 width, f32 height, u32 color) EM_IMPORT(botsDrawRectangle);
void botsDrawCircle(f32 x, f32 y, f32 radius, u32 color) EM_IMPORT(botsDrawCircle);

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

#define PATH_FIND_BUFFER_CAPACITY 4096
static i8 _path_buffer[PATH_FIND_BUFFER_CAPACITY];
static Buffer path_find_buffer = {
	_path_buffer,
	0,
	PATH_FIND_BUFFER_CAPACITY
};

// string must be either null terminated or passed in a string length
void write_string_to_buffer(Buffer* buffer, const i8* string_ptr, const i32 string_length)
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

typedef struct {
    Vec* data;
    i32 size;
    i32 capacity;
} VecArray;

void initVecArray(VecArray* arr, i32 initial_capacity) {
    arr->data = (Vec*)malloc(sizeof(Vec) * initial_capacity);
    arr->size = 0;
    arr->capacity = initial_capacity;
}

void freeVecArray(VecArray* arr) {
    free(arr->data);
    arr->data = NULL;
    arr->size = 0;
    arr->capacity = 0;
}

void pushVec(VecArray* arr, Vec value) {
    if (arr->size >= arr->capacity) {
        i32 new_capacity = arr->capacity * 2;
        Vec* new_data = (Vec*)realloc(arr->data, sizeof(Vec) * new_capacity);
        arr->data = new_data;
        arr->capacity = new_capacity;
    }
    arr->data[arr->size++] = value;
}

void log_message(Buffer *print_buffer, const char *format, ...) {
	reset_buffer(print_buffer);
	va_list args;
	va_start(args, format);
	i32 bytes_written = vsprintf((char *)print_buffer->ptr, format, args);
	my_assert(bytes_written >= 0);
	print_buffer->offset += bytes_written;
	botsLog(print_buffer->ptr, print_buffer->offset);
}

typedef struct {
	u8 team_id;
	Vec base;
	Vec flag;
	u8 is_flag_taken;
	u64 entity_id;
	u32 num_blocks_in_entity;
} Team;

static Team enemy_team = {};
static Team my_team = {}; 

static VecArray my_entity_path = {};
static Vec enemy_position = {};

EMSCRIPTEN_KEEPALIVE void step()
{
	botsDrawText("im with stupid", sizeof("im with stupid"), 100.0f, 100.0f, 20.0f, 0xFF00FF00);
	botsDrawLine(0.0f, 0.0f, 100.0f, 100.0f, 5.0f, 0xffFF0000);
	botsDrawRectangle(50.0f, 200.0f, 100.0f, 100.0f, 0x7f0000FF);
	botsDrawCircle(300.0f, 300.0f, 50.0f, 0xffFFFF00);
	if (step_count == 0) {
		const i32 game_state_buffer_capacity = 64 * 1024;
		game_state_buffer.ptr = (i8*) malloc(game_state_buffer_capacity);
		game_state_buffer.offset = 0;
		game_state_buffer.capacity = game_state_buffer_capacity;

		initVecArray(&my_entity_path, 1024);
	}

	i32 byte_size = botsGetGameState(game_state_buffer.ptr, game_state_buffer.capacity);
	my_assert(byte_size > game_state_buffer.capacity);

	size_t size;
	void * buffer_offset = flatbuffers_read_size_prefix((void*) game_state_buffer.ptr, &size);
	GameStateDelta_table_t game_state = GameStateDelta_as_root(buffer_offset);

	u8 my_id = GameStateDelta_my_id_get(game_state);
	my_team.team_id = my_id;

	Flag_vec_t flags = GameStateDelta_flag_updates_get(game_state);

	size_t num_flags = Flag_vec_len(flags);
	i32 i;
	for (i = 0; i < num_flags; i++) {
		Flag_table_t flag = Flag_vec_at(flags, i);
		u8 owner_id = Flag_owner_id_get(flag);
		f32 x = Flag_x_get(flag);
		f32 y = Flag_y_get(flag);
		if (step_count == 0) {
			if (owner_id == my_id) {
				my_team.team_id = owner_id;
				my_team.base.x = x;
				my_team.base.y = y;
			}  else {
				enemy_team.team_id = owner_id;
				enemy_team.base.x = x;
				enemy_team.base.y = y;
			}
		}
		if (owner_id == my_id) {
			my_team.flag = (Vec) {x, y};
		} else {
			enemy_team.flag = (Vec) {x, y};
		}
	}

	if (step_count == 0) {
		//log_message(&print_buffer, "i love numbers %d %d %d %d", 47, 1337, 911, 101);
		Entity_vec_t new_entities = GameStateDelta_new_entities_get(game_state);
		for (i = 0; i < Entity_vec_len(new_entities); i++) {
			Entity_table_t entity = Entity_vec_at(new_entities, i);
			u8 owner_id = Entity_owner_get(entity);
			u64 id = Entity_id_get(entity);
			Block_vec_t blocks = Entity_blocks_get(entity);
			size_t num_blocks = Block_vec_len(blocks);
			if (owner_id == my_id) {
				my_team.entity_id = id;
				my_team.num_blocks_in_entity = num_blocks;
				log_message(&print_buffer, "my_team.entity_id=%llu, num_blocks=%zu, owner_id=%u, my_id=%u", my_team.entity_id, num_blocks, owner_id, my_id);
			} else if (owner_id > 0) {
				enemy_team.entity_id = id;
				enemy_team.num_blocks_in_entity = num_blocks;
				log_message(&print_buffer, "enemy_team.entity_id=%llu, num_blocks=%zu, owner_id=%u, my_id=%u", enemy_team.entity_id, num_blocks, owner_id, my_id);
			}
		}
	}


	EntityUpdate_vec_t entity_updates = GameStateDelta_entity_updates_get(game_state);
	size_t num_entity_updates = EntityUpdate_vec_len(entity_updates);

	for (i = 0; i < num_entity_updates; i++) {
		EntityUpdate_table_t entity = EntityUpdate_vec_at(entity_updates, i);
		Vec2_struct_t position = EntityUpdate_position_get(entity);
		u64 entity_id = EntityUpdate_id_get(entity);

		//log_message(&print_buffer, "entity update: entity_id=%llu, my_team.entity_id=%llu, enemy_team.entity_id=%llu", entity_id, my_team.entity_id, enemy_team.entity_id);
		if (entity_id == my_team.entity_id) {
			u32 path_byte_size = 0;
			if (enemy_team.is_flag_taken) {
				path_byte_size = botsFindPath(
					Vec2_x_get(position), Vec2_y_get(position), my_team.base.x, my_team.base.y, (void *) path_find_buffer.ptr, path_find_buffer.capacity);
			} else {
				path_byte_size = botsFindPath(
					Vec2_x_get(position), Vec2_y_get(position), enemy_team.flag.x, enemy_team.flag.y, (void *) path_find_buffer.ptr, path_find_buffer.capacity);
			}

			if (path_byte_size > 0) {
				void * buffer_offset = flatbuffers_read_size_prefix((void*) path_find_buffer.ptr, &size);
				Path_table_t path = Path_as_root(buffer_offset);
				Vec2_vec_t path_points = Path_waypoints_get(path);
				size_t num_points = Vec2_vec_len(path_points);

				log_message(&print_buffer, "path_byte_size=%d, waypoints=%zu", path_byte_size, num_points);

				size_t j = 0;
				for (j = num_points-1; j >= 0; j--) {
					Vec2_struct_t waypoint = Vec2_vec_at(path_points, j);
					f32 dx = Vec2_x_get(position) - Vec2_x_get(waypoint);
					f32 dy = Vec2_y_get(position) - Vec2_y_get(waypoint);
					f32 dist = sqrtf(dx * dx + dy * dy);
					if (dist > 100.0f || j == 0) {
						botsMoveEntityToTarget(entity_id, Vec2_x_get(waypoint), Vec2_y_get(waypoint));
						log_message(&print_buffer, "moving to waypoint %zu, dist=%f", j, dist);
						break;
					}
				}
			}

			i32 block_idx = 0;
			for (block_idx = 0; block_idx < my_team.num_blocks_in_entity; block_idx++) {
				botsAimTurret(entity_id, block_idx, enemy_position.x, enemy_position.y);
				botsFireCannon(entity_id, block_idx);
				botsGrabFlag(entity_id);
			}
		}

		if (entity_id == enemy_team.entity_id) {
			enemy_position.x = Vec2_x_get(position);
			enemy_position.y = Vec2_y_get(position);
		}
	}

	reset_buffer(&print_buffer);
	if (step_count > 0 && step_count % 100 == 0) {
		i32 bytes_written = sprintf((char *)print_buffer.ptr, "Hello World. I am C. %d steps. %zu entity updates", step_count, num_entity_updates);
		my_assert(bytes_written >= 0);
		print_buffer.offset += bytes_written;
		botsLog(print_buffer.ptr, print_buffer.offset);
	}

	step_count += 1;
}

