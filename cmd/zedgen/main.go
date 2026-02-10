package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Emyrk/zedgen"
)

func main() {
	var (
		schemaPath = flag.String("schema", "", "Path to SpiceDB schema file (.zed)")
		pkgName    = flag.String("package", "policy", "Go package name for generated code")
		outPath    = flag.String("out", "", "Output file path (default: stdout)")
	)
	flag.Parse()

	if *schemaPath == "" {
		fmt.Fprintln(os.Stderr, "error: -schema flag is required")
		flag.Usage()
		os.Exit(1)
	}

	schema, err := os.ReadFile(*schemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading schema: %v\n", err)
		os.Exit(1)
	}

	output, err := zedgen.Generate(string(schema), zedgen.Options{
		Package:        *pkgName,
		SchemaFileName: *schemaPath,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating code: %v\n", err)
		os.Exit(1)
	}

	if *outPath == "" {
		fmt.Print(output)
	} else {
		if err := os.WriteFile(*outPath, []byte(output), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "wrote %s\n", *outPath)
	}
}
