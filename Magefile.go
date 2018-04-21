// +build mage

/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

func step(message string, args ...interface{}) {
	fmt.Printf("\x1b[33m--- %s\x1b[0m\n", fmt.Sprintf(message, args...))
}

func execute(env map[string]string, args ...string) error {
	step("Executing: %s ...", strings.Join(args, " "))

	_, err := sh.Exec(env, os.Stdout, os.Stderr, args[0], args[1:]...)

	return err
}

// Runs tests.
func Test() error {
	return execute(nil, "go", "test")
}

// Runs tests with coverage.
func Coverage() error {
	return execute(nil, "go", "test", "-coverprofile=coverage.out")
}

// Shows last coverage results.
func ViewCoverage() error {
	return execute(nil, "go", "tool", "cover", "-html=coverage.out")
}

// Verifies the code.
func Lint() error {
	return execute(nil, "go", "vet")
}
