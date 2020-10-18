package go_mod_replace_B

import (
	"fmt"
	x "github.com/alokmenghrajani/go-mod-replace-A"
)

func LibB() string {
	return fmt.Sprintf("LibB: %s\n", x.LibA())
}
