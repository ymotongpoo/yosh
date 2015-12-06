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
	"fmt"
	"os"
)

type builtinFunc func(args []string) ExitStatus

var (
	builtins = map[string]builtinFunc{
		"cd":   cd,
		"ls":   ls,
		"exit": exit,
		"help": help,
	}

	builtinNames = []string{"cd", "ls", "exit", "help"}
)

// cd
func cd(args []string) ExitStatus {
	if len(args) == 1 {
		fmt.Println("cd expects one argument. (target dir)") // TODO(ymotongpoo): change behavior to move to home dir.
		return EXIT_FAILURE
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println(err)
		return EXIT_FAILURE
	}
	return EXIT_SUCCESS
}

// ls
func ls(args []string) ExitStatus {
	var filenames []string
	if len(args) == 1 {
		filenames = []string{"."}
	} else {
		filenames = args[1:]
	}

	status := EXIT_SUCCESS
	for _, f := range filenames {
		file, err := os.Open(f)
		if err != nil {
			status = EXIT_FAILURE
			continue
		}
		fis, err := file.Readdir(-1)
		if err != nil {
			status = EXIT_FAILURE
			continue
		}
		for _, fi := range fis {
			fmt.Printf("%v\t", fi.Name())
		}
	}
	fmt.Println()
	return status
}

// exit
func exit(args []string) ExitStatus {
	return EXIT_SUCCESS
}

// help
func help(args []string) ExitStatus {
	fmt.Println("Yosh: Shell implemented in Go")
	fmt.Println("following commands are builtin functions")
	printBuiltin()
	return EXIT_SUCCESS
}

func printBuiltin() {
	for _, n := range builtinNames {
		fmt.Printf("\t%v\n", n)
	}
}
