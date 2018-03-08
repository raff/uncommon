package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tylertreat/BoomFilters"
)

func error(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	max := flag.Uint("max", 100000, "maximum number of items to compare (lines in first file)")
	fpRate := flag.Float64("fp", 0.01, "false positive rate")
        verbose := flag.Bool("verbose", false, "verbose mode")

	flag.Parse()

	if flag.NArg() != 2 {
		error("usage: uncommon file1 file2")
	}

	var r1, r2 io.Reader

	if f, err := os.Open(flag.Arg(0)); err != nil {
		error(err)
	} else {
		r1 = f
		defer f.Close()
	}

	if flag.Arg(1) == "--" {
		r2 = os.Stdin
	} else if f, err := os.Open(flag.Arg(1)); err != nil {
		error(err)
	} else {
		r2 = f
		defer f.Close()
	}

        if *verbose {
	    fmt.Println("loading", flag.Arg(0))
        }

	cf := boom.NewBloomFilter(*max, *fpRate)
	reader := bufio.NewReaderSize(r1, 128*1024)

	for {
		if line, _, err := reader.ReadLine(); err != nil && err != io.EOF {
			error(flag.Arg(0), err)
		} else if line == nil {
			break
		} else {
			cf.Add(line)
		}
	}

        if *verbose {
	    fmt.Println("loaded", cf.Count(), "entries")
	    fmt.Println("checking", flag.Arg(1))
        }

	reader = bufio.NewReaderSize(r2, 128*1024)

	for {
		if line, _, err := reader.ReadLine(); err != nil && err != io.EOF {
			error(flag.Arg(0), err)
		} else if line == nil {
			break
		} else if cf.Test(line) == false {
			fmt.Println(string(line))
			//fmt.Println("-", string(line))
		} else {
			//fmt.Println("+", string(line))
		}
	}
}
