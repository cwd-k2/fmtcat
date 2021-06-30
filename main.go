package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Cat struct {
	*os.File
}

func (f *Cat) Format(s fmt.State, c rune) {
	switch c {
	case 's':
		io.Copy(s, f)
	}
}

var rep = strings.NewReplacer(
	"\\n", "\n",
	"\\t", "\t",
	"\\a", "\a",
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "fmtcat: not enough arguments")
		os.Exit(1)
	}

	var (
		fmtstr = rep.Replace(os.Args[1])
		cats   = make([]interface{}, 0)
	)

	for _, arg := range os.Args[2:] {
		fp, err := os.Open(arg)
		if err != nil && err != os.ErrNotExist {
			fmt.Fprintf(os.Stderr, "fmtcat: no such file or directory: %s\n", arg)
			cats = append(cats, "")
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "fmtcat: error: %s\n", err.Error())
			os.Exit(1)
		} else {
			defer fp.Close()
			cats = append(cats, &Cat{fp})
		}
	}

	fmt.Fprintf(os.Stdout, fmtstr, cats...)
}
