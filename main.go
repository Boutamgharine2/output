package main

import (
	"flag"
	"fmt"
	"os"

	ascii "ascii/funcs"
)

func main() {
	var filename string
	args := os.Args
	str, Newargs := ascii.CheckArgs(args) // checking args
	style := Newargs[len(Newargs)-1] + ".txt"

	Filestyle, _ := os.Open(style)
	h := ascii.PrintN(Filestyle, str)

	if len(Newargs) == 4 {
		
		flag.StringVar(&filename, "output", "default", "flag str") // give the file parameters
		flag.Parse()
		

		outputFile, err := os.Create(filename)

		outputFile.WriteString(h)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	}
	fmt.Print(h)
}
