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

package main

import (
    `bytes`
    `os`
    `os/exec`
)

const (
    defaultCpp = "/usr/bin/cpp"
)

func preprocess(name string, defs []string) (string, error) {
    var err error
    var cpp string
    var out bytes.Buffer

    /* find the CPP command */
    if cpp = os.Getenv("CPP"); cpp == "" {
        cpp = defaultCpp
    }

    /* command arguments */
    args := []string {
        "-CC",
        "-nostdinc",
    }

    /* build definations */
    for _, def := range defs {
        args = append(args, "-D" + def)
    }

    /* construct the preprocessor command */
    cmd := exec.Command(
        defaultCpp,
        append(args, name)...,
    )

    /* bind stdio */
    cmd.Stdin  = nil
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr

    /* run the preprocessor */
    if err = cmd.Run(); err != nil {
        return "", err
    } else {
        return string(out.Bytes()), nil
    }
}
