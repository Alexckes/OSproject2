package builtins

import (
	"fmt"
	"io"
)

func Eval(w io.Writer, args ...string) error {
	if args[0] == "" {
		return fmt.Errorf("%w: no arguments", ErrInvalidArgCount)
	}
	if args[0] == "help" && len(args) == 1 {
		return Help()
	}
	if args[0] == "echo" && len(args) >= 2 {
		return Echo(args[1])
	}
	if args[0] == "cd" {
		return ChangeDirectory()
	}
	if args[0] == "env" {
		return EnvironmentVariables(w, args[0])
	}
	if args[0] == "times" {
		return Times()
	}
	if args[0] == "pwd" {
		return Pwd()
	}
	return fmt.Errorf("%w: invalid argument", ErrInvalidArgCount)
}
