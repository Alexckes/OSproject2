package builtins

import (
	"fmt"
)

func Help(args ...string) error {
	fmt.Println("The commands are:\ntimes: shows the current time.\ncd: goes to the home directory.\nenv [wwriter]: writes the environment.\nhelp: shows the commands.\neval [arg]: executes the argument command.\npwd: returns the home directory.\necho [arg]: prints out the argument.")
	return fmt.Errorf(" ")
}
