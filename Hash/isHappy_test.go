package Hash

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	happy := IsHappy(19)
	fmt.Println(happy)
}
