// Copyright 2015 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ExitStatus int

const (
	EXIT_SUCCESS ExitStatus = iota
	EXIT_FAILURE
)

type ShellStatus int

const (
	SHELL_RUNNING ShellStatus = iota
	SHELL_EXIT
)

const LineBufSize = 10000

var PROMPT = map[string][]byte{
	"simple": []byte("> "),
	"gopher": []byte("ʕ◔ϖ◔ʔ 三 "),
}

func launch(args []string) ExitStatus {
	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0], []string{}...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return EXIT_FAILURE
	}
	return EXIT_SUCCESS
}

func execute(args []string) ShellStatus {
	if len(args) == 0 {
		return SHELL_RUNNING
	}
	var exitStatus ExitStatus
	if v, ok := builtins[args[0]]; ok {
		exitStatus = v(args)
	} else {
		exitStatus = launch(args)
	}
	_ = exitStatus // TODO(ymotongpoo): need?
	return SHELL_RUNNING
}

func loop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		size := strconv.Itoa(reader.Buffered())
		prompt := append([]byte(size), PROMPT["gopher"]...)
		os.Stdout.Write(prompt)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("yosh:", err)
			continue
		}
		if len(line) == 0 {
			os.Stdout.Write([]byte("hoge"))
			continue
		}
		args := strings.Fields(string(line))
		_ = execute(args)
	}
}

func main() {
	loop()
	os.Exit(int(EXIT_SUCCESS))
}
