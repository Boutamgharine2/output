package ascii

import (
	"fmt"
	"os"
)

func PrintN(inputfile *os.File, v []string) {
	var char []string
	var chars [][]string
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			fmt.Println()
		}
		return
	}

	for _, word := range v {
		r := []rune(word)
		if word == "" {

			fmt.Println()

			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputfile)

			chars = append(chars, char)
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(chars); j++ {
				fmt.Print(chars[j][i])
			}

			fmt.Println()

		}
	}
	defer inputfile.Close()
}
