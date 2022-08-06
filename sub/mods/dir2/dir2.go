package dir2

import (
	"fmt"
	"os"
)

func ShowDir() {
	p := fmt.Println
	theDir, _ := os.Getwd()
	p(theDir)
}