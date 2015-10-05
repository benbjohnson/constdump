package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("packages required")
		os.Exit(1)
	}

	// Loop over directories.
	for _, arg := range flag.Args() {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, arg, nil, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Extract only the non-test packages.
		fmt.Println(arg)
		for pkgname, pkg := range pkgs {
			if strings.HasSuffix(pkgname, "_test") {
				continue
			}

			// Loop over files in package.
			for _, f := range pkg.Files {

				// Loop over top-level objects in the file.
				for name, obj := range f.Scope.Objects {
					// Ignore non-constant objects.
					if obj.Kind != ast.Con {
						continue
					}

					// Print with the package name.
					fmt.Println(pkg.Name + "." + name)
				}
			}
		}
		fmt.Println("")
	}
}
