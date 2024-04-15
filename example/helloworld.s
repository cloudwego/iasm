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

#define STDOUT      $1
#define IMAGE_BASE  0x04000000

#ifdef __Linux__
#define SYS_exit    $1
#define SYS_write   $4
#elif defined(__Darwin__)
#define SYS_exit    $0x02000001
#define SYS_write   $0x02000004
#else
#error Unsupported operating system.
#endif

.org   IMAGE_BASE
.entry start

start:
    movq    STDOUT, %rdi
    leaq    msg(%rip), %rsi
    movq    $13, %rdx
    movq    SYS_write, %rax
    syscall
    xorl    %edi, %edi
    movq    SYS_exit, %rax
    syscall

msg:
    .ascii "hello, world\n"
