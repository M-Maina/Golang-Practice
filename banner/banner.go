package main

import "fmt"

func main() {
	banner("Go", 6)

	s := "Go"
	fmt.Println("len:", len(s))

	for i, r := range s {
		fmt.Println(i, r)
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
