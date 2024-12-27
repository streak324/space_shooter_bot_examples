# Examples
  - go into each example directory for further instructions on compiling each bot.

## WASM Game Host Functions
Below is a list of functions provided by the WASM Game Host that your bot can import. all exist in the `env` namespace

- botsLog(ptr i32, len i32)
  - ptr -> pointer address to the buffer containing the message
  - len -> length of message in bytes. 
  - Use UTF8 encoding.
- getGameState(ptr i32, capacity i32) -> i32
  - ptr -> pointer address to a buffer where the flatbuffer encoded game state will be written to
  - capacity -> how many bytes the buffer can be written to
  - returns number of bytes of the game state. assume error if buffer capacity is smaller
- moveEntityToTarget(entityId u64, x f32, y f32) -> i32
  - entityId -> the id of the entity you want to move
  - x -> a 32 bit floating point in the x dimension you want the entity to be in
  - y -> a 32 bit floating point in the y dimension you want the entity to be in
  - returns id of the input for result inspection in the next time step
- orientTurret(entityId u64, blockIndex i32, rotation f32) -> i32
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the turret you want to orient
  - rotation -> where you want the turret to be rotated to. NOTE: rotation is relative to the entity, not to the world.
  - returns id of the input for result inspection in the next time step
- botsFireCannon(entityId u64, blockIndex i32) -> i32 
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the cannon you want to fire
  - returns id of the input for result inspection in the next time step
- botsLaunchMissiles(entityId u64, blockIndex i32) -> i32 
  - entityId -> the id of the entity you want to command
  - blockIndex -> the block index of the entity containing the missle launcher you want to fire. NOTE: it will fire all filled slots in the missile launcher
  - returns id of the input for result inspection in the next time step
- botsAimTurret(entityId u64, bockIndex i32, x f32, y f32) -> i32
  - returns id of the input for result inspection in the next time step
- botsGrabFlag(entityId u64) -> i32
  - attempt to grab the enemy flag with the owned entity, if close enough
  - returns id of the input for result inspection in the next time step
- botsDropFlag(entityId u64)
  - attempts to drop the enemy flag if the owned entity is carrying the flag
  - returns id of the input for result inspection in the next time step
- botsReturnFlag(entityId u64)
  - attempts to return your flag back to base if the owned entity is close enough. 
  - returns id of the input for result inspection in the next time step
- botsApplyImpulse(entityId u64, block_index: u32, impulse: f32) -> i32
  - attempts to apply impul
- botsFindPath(entityId u64, start_x: f32, start_y: f32, goal_x: f32, goal_y: f32, ptr: u32, capacity: u32) -> u32
  - attempts to find a path from start to goal.
  - writes the path into the buffer `ptr` as a flatbuffer table with a `waypoints` field.
  - returns the byte size of the encoded path. assume the write failed if buffer `capacity` is smaller than the encoded path
- botsDrawText(text_ptr: u32, text_len: u32, x: f32, y: f32, size: f32, color: u32)
- botsDrawRect(x: f32, y: f32, width: f32, height: f32, color: u32)
- botsDrawLine(x1: f32, y1: f32, x2: f32, y2: f32, thickness: f32, color: u32)
- botsDrawCircle(x: f32, y: f32, radius: f32, color: u32)


## Notes
- Adjust your wasm toolchains to avoid generating imports for "GOT.mem", "GOT.func", "__stack_pointer", "__memory_base", "__table_base ". Those are not well supported in the wasm host. Use the toolchain configuration in the examples
- color parameters are in hexadecimal AARRGGBB format.