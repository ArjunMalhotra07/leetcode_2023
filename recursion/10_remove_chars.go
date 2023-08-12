package recursion

import (
	"fmt"
	"strings"
)

func RemoveCharacter(test, removeCharacter string) string {
	if len(test) == 0 {
		return ""
	}
	if string(test[0]) != removeCharacter {
		return string(test[0]) + RemoveCharacter(string(test[1:]), removeCharacter)
	} else {
		return RemoveCharacter(string(test[1:]), removeCharacter)
	}
}

func RemoveWord(inputString, wordToRemove string) string {
	if len(inputString) == 0 {
		return ""
	}
	if strings.HasPrefix(inputString, wordToRemove) {
		return RemoveWord(string(inputString[len(wordToRemove):]), wordToRemove)
	} else {
		return string(inputString[0]) + RemoveWord(string(inputString[1:]), wordToRemove)
	}
}
func RemoveAppWhenNotApple(inputString, wordToRemove string) string {
	if len(inputString) == 0 {
		return ""
	}
	if strings.HasPrefix(inputString, wordToRemove) && !strings.HasPrefix(inputString, "Apple") {
		return RemoveAppWhenNotApple(string(inputString[len(wordToRemove):]), wordToRemove)
	} else {
		return string(inputString[0]) + RemoveAppWhenNotApple(string(inputString[1:]), wordToRemove)
	}
}

func RemoveCharacterHelper() {
	fmt.Println(RemoveCharacter("ABCAfBghjaA", "B"))
	fmt.Println(RemoveAppWhenNotApple("zyxAppleiAppj", "App"))

}
