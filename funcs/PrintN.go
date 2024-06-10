package ascii

import (
	"os"
)

func PrintN(inputfile *os.File, v []string) string{
	var char []string
	var chars [][]string
	var n string
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			n+="\n"
		}
		
	}

	for _, word := range v {
		r := []rune(word)
		if word == "" {

			n+="\n"

			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputfile)

			chars = append(chars, char)
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(chars); j++ {
				n+=chars[j][i]
			}

			n+="\n"

		}
	}
	defer inputfile.Close()
	return n
}
