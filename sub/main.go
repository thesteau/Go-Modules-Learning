package main

import (
	"fmt"
	d1 "samp/dir1"
	d2 "samp/dir2"
	d3 "samp/dir3"
)

func display() {
	fmt.Println("test")

	d1.ShowQ()
	d2.ShowDir()
	str := d3.UserInput()

	fmt.Println("You have entered:", str)
}

func main() {
	display()
}