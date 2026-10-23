package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func printRunes(ar []rune) {
	for _, v := range ar {
		fmt.Print(v)
	}
}

var endOfLineRune = map[rune]bool{
	'\n': true,
	'\r': true,
}

func readFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	runes := []rune{}
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r != '\n' {
			runes = append(runes, r)
		} else {
			// print the rune slice backwards
			for _, r := range slices.Backward(runes) {
				fmt.Print(string(r))
			}

			// clearing and new space added
			runes = []rune{}
			fmt.Println()
		}
	}
}

func readFromFile(file string) bool {

	dat, err := os.ReadFile(file)

	if program := os.Args[0]; err != nil {
		fmt.Println(program+":", file+":", "No such file or directory")
		return false
	}

	text := string(dat)

	for _, r := range slices.Backward([]rune(text)) {
		fmt.Printf("%q", r)
	}
	fmt.Println()
	return true
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		readFromStdin()

	} else {
		success := true
		for _, file := range args {
			ok := readFromFile(file)
			if ok != false {
				success = false
			}
		}
		if success != true {
			os.Exit(1)
		}
	}
}
