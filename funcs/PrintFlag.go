package ascii

import (
	"os"
)

func PrintFlag(inputFile, outputFile *os.File, v []string) {
	var char []string
	var chars [][]string
	n := ""
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			n+="\n"
		}
		return
	}

	for _, word := range v {
		r := []rune(word)
		if word == "" {

			n+="\n"

			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputFile)

			chars = append(chars, char)
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(chars); j++ {
				n+=chars[j][i]
			}

			n+="\n"

		}
	}
	defer inputFile.Close()
}
