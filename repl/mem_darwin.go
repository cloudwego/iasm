package repl

import (
    `syscall`
)

const (
    _MMAP_FLAGS = syscall.MAP_JIT | syscall.MAP_ANON | syscall.MAP_PRIVATE
)
