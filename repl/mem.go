//
// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package repl

import (
    `os`
    `reflect`
    `syscall`
    `unsafe`
)

type _Memory struct {
    size uint64
    addr uintptr
}

func (self _Memory) p() (v unsafe.Pointer) {
    *(*uintptr)(unsafe.Pointer(&v)) = self.addr
    return
}

func (self _Memory) buf() (v []byte) {
    p := (*reflect.SliceHeader)(unsafe.Pointer(&v))
    p.Cap = int(self.size)
    p.Len = int(self.size)
    p.Data = self.addr
    return
}

const (
    _AP  = syscall.MAP_ANON  | syscall.MAP_PRIVATE
    _RWX = syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC
)

var (
    _PageSize = uint64(os.Getpagesize())
)

func mmap(nb uint64) (_Memory, error) {
    nb = (((nb - 1) / _PageSize) + 1) * _PageSize
    mm, _, err := syscall.RawSyscall6(syscall.SYS_MMAP, 0, uintptr(nb), _RWX, _AP, 0, 0)

    /* check for errors */
    if err != 0 {
        return _Memory{}, err
    } else {
        return _Memory { addr: mm, size: nb }, nil
    }
}

func munmap(m _Memory) error {
    if _, _, err := syscall.RawSyscall(syscall.SYS_MUNMAP, m.addr, uintptr(m.size), 0); err != 0 {
        return err
    } else {
        return nil
    }
}
