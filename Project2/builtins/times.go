package builtins

import (
	"time"
	"fmt"
)



func Times(args ...string) error {
	fmt.Println(""+time.Now().String());
	return fmt.Errorf("");
}