// csvv - A simple CSV extractor.
// Author: Takahiro Yoshihara <tacahiroy@gmail.com>
// License: Modified BSD License

// Copyright © 2016 Takahiro Yoshihara
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
// 1. Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
// notice, this list of conditions and the following disclaimer in the
// documentation and/or other materials provided with the distribution.
// 3. Neither the name of the organization nor the
// names of its contributors may be used to endorse or promote products
// derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY  "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL  BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// Version represents version of the app
const Version = "0.0.3"

var (
	useTab      bool
	listHeaders bool
	delimiter   string
)

func usage() {
	var app = path.Base(os.Args[0])

	fmt.Printf(`%s version %s

  %s [OPTIONS] /path/to/input-file field1[,field2...]

  Options:
  -t Use Tab (0x09) for the field separator
  -l List header names
	`, app, Version, app)
	os.Exit(1)
}

func printLine(cols []string) {
	fmt.Println(strings.Join(cols, delimiter))
}

func init() {
	flag.BoolVar(&useTab, "t", false, "")
	flag.BoolVar(&listHeaders, "l", false, "")
	flag.Parse()

	if useTab {
		delimiter = "\t"
	} else {
		delimiter = ","
	}
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	csvfile, err := os.Open(flag.Arg(0))
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
	var header map[string]int = map[string]int{}
	for index, name := range row {
		header[name] = index
	}

	if listHeaders {
		for header := range header {
			fmt.Println(header)
		}
		os.Exit(0)
	}

	// Parse 2nd argument to determine which columns need to be got
	var cols []string
	for _, c := range strings.Split(flag.Arg(1), ",") {
		if _, ok := header[c]; ok {
			cols = append(cols, c)
		}
	}

	// header
	printLine(cols)

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
			if _, ok := header[col]; ok {
				line = append(line, rec[header[col]])
			}
		}

		printLine(line)
	}
}
