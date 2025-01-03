#ifndef PATH_READER_H
#define PATH_READER_H

/* Generated by flatcc 0.6.2 FlatBuffers schema compiler for C by dvide.com */

#ifndef FLATBUFFERS_COMMON_READER_H
#include "flatbuffers_common_reader.h"
#endif
#include "flatcc/flatcc_flatbuffers.h"
#ifndef __alignas_is_defined
#include <stdalign.h>
#endif
#include "flatcc/flatcc_prologue.h"
#ifndef flatbuffers_identifier
#define flatbuffers_identifier 0
#endif
#ifndef flatbuffers_extension
#define flatbuffers_extension "bin"
#endif

typedef struct Vec2 Vec2_t;
typedef const Vec2_t *Vec2_struct_t;
typedef Vec2_t *Vec2_mutable_struct_t;
typedef const Vec2_t *Vec2_vec_t;
typedef Vec2_t *Vec2_mutable_vec_t;

typedef const struct Path_table *Path_table_t;
typedef struct Path_table *Path_mutable_table_t;
typedef const flatbuffers_uoffset_t *Path_vec_t;
typedef flatbuffers_uoffset_t *Path_mutable_vec_t;
#ifndef Vec2_file_identifier
#define Vec2_file_identifier 0
#endif
/* deprecated, use Vec2_file_identifier */
#ifndef Vec2_identifier
#define Vec2_identifier 0
#endif
#define Vec2_type_hash ((flatbuffers_thash_t)0x2c3c5815)
#define Vec2_type_identifier "\x15\x58\x3c\x2c"
#ifndef Vec2_file_extension
#define Vec2_file_extension "bin"
#endif
#ifndef Path_file_identifier
#define Path_file_identifier 0
#endif
/* deprecated, use Path_file_identifier */
#ifndef Path_identifier
#define Path_identifier 0
#endif
#define Path_type_hash ((flatbuffers_thash_t)0xeb66e456)
#define Path_type_identifier "\x56\xe4\x66\xeb"
#ifndef Path_file_extension
#define Path_file_extension "bin"
#endif


struct Vec2 {
    alignas(4) float x;
    alignas(4) float y;
};
static_assert(sizeof(Vec2_t) == 8, "struct size mismatch");

static inline const Vec2_t *Vec2__const_ptr_add(const Vec2_t *p, size_t i) { return p + i; }
static inline Vec2_t *Vec2__ptr_add(Vec2_t *p, size_t i) { return p + i; }
static inline Vec2_struct_t Vec2_vec_at(Vec2_vec_t vec, size_t i)
__flatbuffers_struct_vec_at(vec, i)
static inline size_t Vec2__size(void) { return 8; }
static inline size_t Vec2_vec_len(Vec2_vec_t vec)
__flatbuffers_vec_len(vec)
__flatbuffers_struct_as_root(Vec2)

__flatbuffers_define_struct_scalar_field(Vec2, x, flatbuffers_float, float)
__flatbuffers_define_struct_scalar_field(Vec2, y, flatbuffers_float, float)


struct Path_table { uint8_t unused__; };

static inline size_t Path_vec_len(Path_vec_t vec)
__flatbuffers_vec_len(vec)
static inline Path_table_t Path_vec_at(Path_vec_t vec, size_t i)
__flatbuffers_offset_vec_at(Path_table_t, vec, i, 0)
__flatbuffers_table_as_root(Path)

__flatbuffers_define_vector_field(0, Path, waypoints, Vec2_vec_t, 0)


#include "flatcc/flatcc_epilogue.h"
#endif /* PATH_READER_H */
