package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flagPrintHelp = flag.Bool("help", false, "print extended usage info")
	flagPrintInfo = flag.Bool("i", false, "print color codes info and exit")
	flagReadFile  = flag.String("s", "", "read from file")

	colorNames      = []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	formattingNames = []string{"reset", "bold", "dim", "italic", "underline"}
)

func Read(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	err := scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func PrintFormattedLine(f string, c int, arr []string) {
	fmt.Printf("  \033[%sm%s%s%d\033[0m\n",
		f,
		arr[c],
		strings.Repeat(" ", 10-len(arr[c])),
		c)
}

func PrintInfo() {
	fmt.Println("colors:")
	for i := 0; i <= 7; i++ {
		if i == 7 {
			PrintFormattedLine(fmt.Sprintf("4%d;30", i), i, colorNames)
			continue
		}
		PrintFormattedLine(fmt.Sprintf("4%d", i), i, colorNames)
	}
	fmt.Println("\nformatting:")
	for i := 0; i <= 3; i++ {
		PrintFormattedLine(fmt.Sprintf("%d", i), i, formattingNames)
	}
}

func PrintHelp() {
	fmt.Println(`golored: color any command's output

Usage: golored [-i] [-help] [-s filename] options...
  -i         print color codes info and exit
  -help      print this help
  -s         read from file (by default stdin is read)

Options define how the text is formated:
  f:color    set foreground color
  b:color    set background color
  bold       bold text
  underline  underlined text
  dim        dimmed text (not supported by all terminals)
  italic     italic text (not supported by all terminals)

Any number of options can passed. bold, dim and italic are combined, in case
of multiple colors the last value overrides the previous ones.

List of colors:
  black, red, green, yellow, blue, magenta, cyan, white `)
}

func GetColorCode(color string) int {
	for i, c := range colorNames {
		if color == c {
			return i
		}
	}
	return 0
}

func GetFormattingCode(option string) int {
	for i, o := range formattingNames {
		if option == o {
			return i
		}
	}
	return 0
}

func main() {
	flag.Parse()

	if *flagPrintInfo {
		PrintInfo()
		return
	}

	if *flagPrintHelp {
		PrintHelp()
		return
	}

	for _, o := range flag.Args() {
		if o[:2] == "f:" {
			fmt.Printf("\033[3%dm", GetColorCode(o[2:]))
			continue
		}
		if o[:2] == "b:" {
			fmt.Printf("\033[4%dm", GetColorCode(o[2:]))
			continue
		}
		fmt.Printf("\033[%dm", GetFormattingCode(o))
	}

	if *flagReadFile == "" {
		Read(os.Stdin)
	} else {
		f, err := os.Open(*flagReadFile)
		if err != nil {
			os.Exit(1)
		}
		Read(f)
	}

	fmt.Print("\033[0m")
}
