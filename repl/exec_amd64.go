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

type _ExecutorAMD64 struct {
    after  _RegFileAMD64
    before _RegFileAMD64
}

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func execaddr(addr uintptr, before *_RegFileAMD64, after *_RegFileAMD64)

func (self *_ExecutorAMD64) Execute(addr uintptr) (_RegFile, _RegFile, error) {
    execaddr(addr, &self.before, &self.after)
    return &self.before, &self.after, nil
}

func init() {
    _exec = new(_ExecutorAMD64)
}
