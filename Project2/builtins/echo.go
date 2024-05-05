package builtins

import (
	"fmt"
)

func Echo(args ...string) error {
	if args[0] == "" {
		return fmt.Errorf("%w: no arguments", ErrInvalidArgCount)
	}
	fmt.Println("" + args[0])
	return fmt.Errorf(" ")
}
