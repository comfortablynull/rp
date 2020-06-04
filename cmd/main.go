package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/comfortablynull/rp/internal"
)

func compile(p string, template *template.Template) {
	f, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := internal.Compile(f, template, "Seed"); err != nil {
		log.Fatal(err)
	}
	if pabs, err := filepath.Abs(p); err == nil {
		fmt.Println("Wrote:", pabs)
	}
}

func main() {
	p := "./"
	temps := map[string]*template.Template{
		"global.go":    internal.Global,
		"decorator.go": internal.Decorator,
		"rand.go":      internal.Interface,
	}
	for k, v := range temps {
		compile(path.Join(p, k), v)
	}
	bm := path.Join(p, "benchmark_test.go")
	f, err := os.Create(bm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := internal.CompileBench(f, "Seed", "Shuffle", "Read"); err != nil {
		log.Fatal(err)
	}
}
