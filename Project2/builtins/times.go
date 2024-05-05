package builtins

import (
	"fmt"
	"time"
)

func Times(args ...string) error {
	fmt.Println("The time is: " + time.Now().String())
	return fmt.Errorf(" ")
}
