package StringAlgo

import (
	"fmt"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	str := " Sakura  Hello  "
	words := reverseWords(str)
	fmt.Println(words)

}

func Test_reverseWords2(t *testing.T) {
	str := " Sakura  Hello  "
	words2 := reverseWords2(str)
	fmt.Println(words2)
}
