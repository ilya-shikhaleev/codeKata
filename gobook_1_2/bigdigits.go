package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <number>\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	number := os.Args[1]
	log.Println(number)

	for row := range bigDigits[0] {
		line := ""
		for _, currChar := range number {
			digit, err := strconv.Atoi(string(currChar))
			if err != nil {
				log.Fatal("Invalid number ", number, ", char is not digit ", currChar)
			}

			line += bigDigits[digit][row] + " "
		}
		fmt.Println(line)
	}
}

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		"   1",
		" 1 1",
		"1  1",
		"   1",
		"   1",
		"   1",
	},
	{
		" 222 ",
		"2   2",
		"    2",
		"  22 ",
		" 2   ",
		"22222",
	},
	{
		" 333 ",
		"3   3",
		"   3 ",
		"    3",
		"3   3",
		" 333 ",
	},
	{
		"4   4",
		"4   4",
		"4   4",
		"44444",
		"    4",
		"    4",
	},
	{
		"55555",
		"5    ",
		"5555 ",
		"    5",
		"    5",
		"5555 ",
	},
	{
		"  666 ",
		" 6    ",
		"66666 ",
		"6    6",
		"6    6",
		" 6666 ",
	},
	{
		"777777",
		"    7 ",
		"   7  ",
		"  7   ",
		" 7    ",
		"7     ",
	},
	{
		" 8888 ",
		"8    8",
		" 8888 ",
		"8    8",
		"8    8",
		" 8888 ",
	},
	{
		" 9999 ",
		"9    9",
		" 99999",
		"     9",
		"    9 ",
		"  99  ",
	},
}
