package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/oliveagle/jsonpath"
)

const name = "jsonuniq"

const version = "0.0.1"

var revision = "HEAD"

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func run() error {
	var jp string

	flag.StringVar(&jp, "p", "$.id", "jsonpath to the value for compareing column")
	flag.Parse()

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	seen := map[any]struct{}{}
	for {
		var v any
		err := dec.Decode(&v)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		v1, err := jsonpath.JsonPathLookup(v, jp)
		fatalIf(err)
		if _, ok := seen[v1]; !ok {
			enc.Encode(v)
			seen[v1] = struct{}{}
		}
	}
	return nil
}

func main() {
	fatalIf(run())
}
