package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: csvv CSV-FILE col1[,col2,col3 ...]")
		return
	}

	csvfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()
	reader := csv.NewReader(csvfile)

	// Get all headers. assuming the first line is header line
	row, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}
	var h map[string]int = map[string]int{}
	for index, name := range row {
		h[name] = index
	}

	// Parse 2nd argument to determine which columns need to be got
	var cols []string
	for _, c := range strings.Split(os.Args[2], ",") {
		if _, ok := h[c]; ok {
			cols = append(cols, c)
		}
	}

	// header
	PrintCSV(cols)

	// Parse body
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		var line []string
		for _, col := range cols {
			if _, ok := h[col]; ok {
				line = append(line, rec[h[col]])
			}
		}
		PrintCSV(line)
	}
}

func Pos(sl []string, v string) int {
	for index, value := range sl {
		if value == v {
			return index
		}
	}
	return -1
}

func PrintCSV(cols []string) {
	fmt.Println(strings.Join(cols, ","))
}
