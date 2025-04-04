package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	VERSION    string = "unknown"
	COMMIT_SHA string = "unknown"

	flagPrintHelp    = flag.Bool("help", false, "print extended usage info")
	flagPrintVersion = flag.Bool("version", false, "print program version and exit")
	flagPrintInfo    = flag.Bool("i", false, "print color codes info and exit")
	flagPrintExtInfo = flag.Bool("ii", false, "print extended color codes info and exit")
	flagReadFile     = flag.String("s", "", "read from file")

	bgFgCodes       = map[rune]int{'f': 3, 'b': 4}
	colorNames      = []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	formattingNames = []string{"reset", "bold", "dim", "italic", "underline"}
)

// read and output contents of a file
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

// get ansi code for color name
func GetColorCode(color string) int {
	for i, c := range colorNames {
		if color == c {
			return i
		}
	}
	return 0
}

// get ansi code for formatting option
func GetFormattingCode(option string) int {
	for i, o := range formattingNames {
		if option == o {
			return i
		}
	}
	return 0
}

func main() {
	flag.Usage = PrintHelp
	flag.Parse()

	// version, help and info
	if *flagPrintVersion {
		PrintVersion()
		return
	}
	if *flagPrintInfo {
		PrintInfo()
		return
	}
	if *flagPrintExtInfo {
		PrintInfo()
		PrintExtInfo()
		return
	}
	if *flagPrintHelp {
		PrintHelp()
		return
	}

	// parse arguments and set color
	for _, o := range flag.Args() {
		if o[:2] == "f:" || o[:2] == "b:" {
			num, err := strconv.Atoi(o[2:])
			if err == nil {
				fmt.Printf("\033[%d8;5;%dm", bgFgCodes[rune(o[0])], num)
				continue
			}
			fmt.Printf("\033[%d%dm", bgFgCodes[rune(o[0])], GetColorCode(o[2:]))
			continue
		}
		fmt.Printf("\033[%dm", GetFormattingCode(o))
	}

	// read and write text
	if *flagReadFile == "" {
		Read(os.Stdin)
	} else {
		f, err := os.Open(*flagReadFile)
		if err != nil {
			os.Exit(1)
		}
		Read(f)
	}

	// reset at the end
	fmt.Print("\033[0m")
}
