package dir1

import (
	"fmt"
	"rsc.io/quote"
)

func ShowQ() {
	p := fmt.Println
	p(quote.Go())
}