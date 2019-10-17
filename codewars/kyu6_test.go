package codewars

import (
	"fmt"
	"testing"
)

func TestDecodeMorse(t *testing.T)  {
	morse := ".... . -.--   .--- ..- -.. ."
	fmt.Println(DecodeMorse(morse))
}
