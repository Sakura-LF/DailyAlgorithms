package Hash

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	anagram := isAnagram("anagram", "nagaram")
	fmt.Println(anagram)

	anagram2 := isAnagram("a", "ab")
	fmt.Println(anagram2)
}
