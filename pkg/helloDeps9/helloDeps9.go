package helloDeps9

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps9.PrintPhrase")
	return phrase
}
