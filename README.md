# Examples
  - go into each example directory for further instructions on compiling each bot.

## WASM Game Host Functions
Below is a list of functions provided by the WASM Game Host that your bot can import:

- botsLog(ptr i32, len i32)
  - ptr -> pointer address to the buffer containing the message
  - len -> length of message in bytes. 
  - Use UTF8 encoding.
- getGameState(ptr i32, capacity i32) -> i32
  - ptr -> pointer address to a buffer where the flatbuffer encoded game state will be written to
  - capacity -> how many bytes the buffer can be written to
  - returns number of bytes written. assume error if buffer capacity is smaller
- botsGetMyOwnerId() -> i32
  - returns the owner id of your controllable entities. may be useful for identifying friend and foe.
- moveEntityToTarget(entityId u64, x f32, y f32) -> i32
  - entityId -> the id of the entity you want to move
  - x -> a 32 bit floating point in the x dimension you want the entity to be in
  - y -> a 32 bit floating point in the y dimension you want the entity to be in
  - returns non-zero on validation error
- orientTurret(entityId u64, blockIndex i32, rotation f32) -> i32
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the turret you want to orient
  - rotation -> where you want the turret to be rotated to. NOTE: rotation is relative to the entity, not to the world.
  - returns non-zero on validation error
- botsFireCannon(entityId u64, blockIndex i32) -> i32 
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the cannon you want to fire
  - returns non-zero on validation error
- botsLaunchMissiles(entityId u64, blockIndex i32) -> i32 
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the missle launcher you want to fire. NOTE: it will fire all filled slots in the missile launcher
  - returns non-zero on validation error
- botsAimTurret(entityId u64, bockIndex i32, x f32, y f32) -> i32
  - returns non-zero on validation error

## Notes
- Adjust your wasm toolchains to avoid generating imports for "GOT.mem", "GOT.func", "__stack_pointer", "__memory_base", "__table_base ". Those are not well supported in the wasm host. Use the toolchain configuration in the examples