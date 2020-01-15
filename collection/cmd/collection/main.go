package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"strings"
)

var (
	t      = flag.String("type", "", "type name; must be set")
	output = flag.String("output", "", "output file name; default srcdir/<type>.go")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of collection:\n")
	fmt.Fprintf(os.Stderr, "\tcollection [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tcollection [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

const f = `func %[2]s(length int, do func(i int) %[1]s, funcs ...CollecitonFunc) []%[1]s {
	var o CollecitonOptions
	for _,f := range funcs {
		f(&o)
	}
	if o.judge == nil {
		o.judge = func(i int) bool { return true }
	}
	result := make([]%[1]s, length)
	filter.Filter(length, o.judge, func(i int) { result[i] = do(i) })
	return result
}`

type Generator struct {
	buf bytes.Buffer // Accumulated output.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("collection: ")
	flag.Usage = Usage
	flag.Parse()

	if len(*t) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	g := Generator{}
	g.Printf("// Code generated by \"collection %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package collection")
	g.Printf("\n")
	g.Printf("import \"github.com/t3reezhou/golang_util/filter\"\n") // Used by all methods.
	bs := []byte(*t)
	bs[0] = bytes.ToUpper(bs[:1])[0]
	g.Printf(f, *t, string(bs))

	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s.go", *t)
		outputName = filepath.Join("", strings.ToLower(baseName))
	}
	err := ioutil.WriteFile(outputName, g.format(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
