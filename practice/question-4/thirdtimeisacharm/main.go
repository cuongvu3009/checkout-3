// https://github.com/01-edu/public/tree/master/subjects/thirdtimeisacharm

package main

import (
	"fmt"
)

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return ""
	}

	var answer string
	for i := 2; i < len(str); i = i + 3 {
		answer = answer + string(str[i])
	}

	return answer + "\n"
}

// Output:
// 369
// b e
