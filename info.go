package main

import (
	"fmt"
	"strings"
)

// formatted line for short info
func PrintFormattedInfoLine(f string, c int, arr []string) {
	fmt.Printf("  \033[%sm%s%s%d\033[0m\n",
		f,
		arr[c],
		strings.Repeat(" ", 10-len(arr[c])),
		c)
}

// short info with main colors
func PrintInfo() {
	fmt.Println("colors:")
	for i := 0; i <= 7; i++ {
		if i == 7 {
			PrintFormattedInfoLine(fmt.Sprintf("4%d;30", i), i, colorNames)
			continue
		}
		PrintFormattedInfoLine(fmt.Sprintf("4%d", i), i, colorNames)
	}
	fmt.Println("\nformatting:")
	for i := 0; i <= 3; i++ {
		PrintFormattedInfoLine(fmt.Sprintf("%d", i), i, formattingNames)
	}
}

// formatted number for short info
func PrintFormattedExtInfoNumber(pre string, color int) {
	fmt.Printf("\033[%s48;5;%dm%s%d\033[0m ",
		pre,
		color,
		strings.Repeat(" ", 3-len(fmt.Sprintf("%d", color))),
		color)
}

// print all 8-bit colors
func PrintExtInfo() {
	fmt.Println("\n8-bit colors:")

	// "normal" colors
	fmt.Print("  ")
	for i := 0; i <= 15; i++ {
		PrintFormattedExtInfoNumber("30;", i)
	}

	// 232-255 (gray-scale)
	for i := 0; i <= 1; i++ {
		fmt.Print("\n  ")
		for j := 0; j <= 11; j++ {
			PrintFormattedExtInfoNumber(strings.Repeat("30;", i), 232+(12*i)+j)
		}
	}

	// 16-231 (colorful, dark and bright)
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Print("\n  ")
			for k := 0; k <= 2; k++ {
				for l := 0; l <= 5; l++ {
					PrintFormattedExtInfoNumber(
						strings.Repeat("30;", i),
						16+(18*i)+(36*j)+(6*k)+l)
				}
				fmt.Print(" ")
			}
		}
	}

	fmt.Println()
}

func PrintHelp() {
	fmt.Println(`golored: color any command's output

Usage: golored [-i] [-ii] [-help] [-s filename] options...
  -i         print color codes info and exit
  -ii        print extended color codes info (8-bit) and exit
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
  black, red, green, yellow, blue, magenta, cyan, white
  0-255 (8-bit colors)`)
}
