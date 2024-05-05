package builtins

import (
	"fmt"
	"os"
)

var (
	Home, _ = os.UserHomeDir()
)

func Pwd(args ...string) error {
	if Home == "" {
		return fmt.Errorf("%w: no homedir found", ErrInvalidArgCount)
	}
	fmt.Println("" + Home)
	return fmt.Errorf(" ")
}
