package builtins

import (
	"fmt"
)

func Help(args ...string) error {
	fmt.Println("The commands are:\ntimes: shows the current time.\ncd: go to home directory.\nenv:")
	return fmt.Errorf(" ")
}
