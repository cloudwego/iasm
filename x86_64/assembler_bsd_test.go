//go:build darwin
// +build darwin

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

package x86_64

import (
    `fmt`
    `os`
    `os/exec`
    `testing`

    `github.com/cloudwego/iasm/obj`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`
)

func TestAssembler_Assemble(t *testing.T) {
    p := new(Assembler)
    e := p.Assemble(`
.org 0x08000000
.entry start

msg:
    .ascii "hello, world\n"

start:
    movl    $1, %edi
    leaq    msg(%rip), %rsi
    movl    $13, %edx
    movl    $0x02000004, %eax
    syscall
    xorl    %edi, %edi
    movl    $0x02000001, %eax
    syscall
`)
    require.NoError(t, e)
    code := p.Code()
    base := p.Base()
    entry := p.Entry()
    spew.Dump(code)
    fmt.Printf("Image Base  : %#x\n", base)
    fmt.Printf("Entry Point : %#x\n", entry)
    fp, err := os.CreateTemp("", "macho_out-")
    require.NoError(t, err)
    err = obj.MachO.Write(fp, code, uint64(base), uint64(entry))
    require.NoError(t, err)
    err = fp.Close()
    require.NoError(t, err)
    err = os.Chmod(fp.Name(), 0755)
    require.NoError(t, err)
    println("Saved to", fp.Name())
    out, err := exec.Command(fp.Name()).Output()
    require.NoError(t, err)
    spew.Dump(out)
    require.Equal(t, []byte("hello, world\n"), out)
    err = os.Remove(fp.Name())
    require.NoError(t, err)
}
