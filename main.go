package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	backgroundColor = flag.String("b", "", "background color")
	foregroundColor = flag.String("f", "", "foreground color")
	formatting      = flag.String("F", "", "formatting options")
	readFile        = flag.String("s", "", "read from file")
	printInfo       = flag.Bool("i", false, "print color codes info and exit")

	colorNames      = []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	formattingNames = []string{"bold", "dim", "italic", "underline"}
)

func main() {
	flag.Parse()

	if *printInfo {
		fmt.Println("colors:")
		for i := 0; i <= 7; i++ {
			if i == 7 {
				fmt.Printf("  \033[4%d;30m%s%s%d\033[0m\n",
					i,
					colorNames[i],
					strings.Repeat(" ", 10-len(colorNames[i])),
					i)
				continue
			}
			fmt.Printf("  \033[4%dm%s%s%d\033[0m\n",
				i,
				colorNames[i],
				strings.Repeat(" ", 10-len(colorNames[i])),
				i)
		}
		fmt.Println("\nformatting:")
		for i := 0; i <= 3; i++ {
			fmt.Printf("  \033[%dm%s%s%d\033[0m\n",
				i+1,
				formattingNames[i],
				strings.Repeat(" ", 10-len(formattingNames[i])),
				i+1)
		}
		return
	}
}
