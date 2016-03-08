package main

import (
	"bufio"
	"fmt"
	flag "github.com/ogier/pflag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var bar bool

func init() {
	flag.BoolVarP(&bar, "bar", "b", false, "Show bar lines")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <number>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func writeDigits(number string, buf io.Writer) (countColumns int) {
	var line string
	for row := range bigDigits[0] {
		line = ""
		for _, currChar := range number {
			digit, err := strconv.Atoi(string(currChar))
			if err != nil {
				log.Fatal("Invalid number ", number, ", char is not digit ", currChar)
			}

			line += bigDigits[digit][row] + " "
		}
		line = line[:len(line)-1]
		fmt.Fprintln(buf, line)
	}
	countColumns = len(line)

	return
}

func main() {
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)
	number := flag.Arg(0)

	if len(number) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	countColumns := writeDigits(number, writer)

	if bar {
		fmt.Println(strings.Repeat("*", countColumns))
	}
	writer.Flush()
	if bar {
		fmt.Println(strings.Repeat("*", countColumns))
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
		"4   4 ",
		"4   4 ",
		"4   4 ",
		"444444",
		"    4 ",
		"    4 ",
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
