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
    `errors`

    `github.com/cloudwego/iasm/x86_64`
)

type _IASMArchSpecific struct {
    ps x86_64.Parser
}

func (self *_IASMArchSpecific) doasm(addr uintptr, line string) ([]byte, error) {
    var err error
    var asm x86_64.Assembler
    var ast *x86_64.ParsedLine

    /* parse the line */
    if ast, err = self.ps.Feed(line); err != nil {
        return nil, err
    }

    /* interactive shell does not support labels */
    if ast.Kind == x86_64.LineLabel {
        return nil, errors.New("interactive shell does not support labels")
    }

    /* assemble the line */
    if err = asm.WithBase(addr).Assemble(line); err != nil {
        return nil, err
    } else {
        return asm.Code(), nil
    }
}
