// The entry file of your WebAssembly module.

@external("env", "botsLog")
declare function botsLog(ptr: u32, size: u32): void

export function myAbort(message: usize, fileName: usize, line: u32, column: u32): void {
  unreachable();
}

export function step(): void {
  const buffer = String.UTF8.encode("assemblyscript reporting for duty");
  botsLog(changetype<u32>(buffer), buffer.byteLength);
}
