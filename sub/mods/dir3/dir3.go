package dir3

import (
	"fmt"
)

func UserInput() string {
	p := fmt.Print
	p("Enter some information: ")

	var str string
	fmt.Scanf("%s", &str)

	return str
}