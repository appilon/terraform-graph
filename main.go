// terraform-graph opens a d3 powered collapsible graph
// Usage:
//
//    # Open a file in a browser window
//    terraform-graph $FILE
//
//    # Open the contents of stdin in a browser window
//    cat $SOMEFILE | terraform-graph
package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/browser"
)

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n  %s [file]\n", os.Args[0])
	flag.PrintDefaults()
}

func must(errs ...error) {
	for _, err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getData() []byte {
	args := flag.Args()
	var data []byte
	var err error

	switch len(args) {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
	case 1:
		data, err = ioutil.ReadFile(args[0])
	default:
		usage()
		os.Exit(1)
	}
	must(err)

	return data
}

func main() {
	// TODO use html/template
	graph := html.EscapeString(string(getData()))
	// dagre-d3 doesn't like shape=box
	// https://github.com/dagrejs/dagre-d3/issues/131
	graph = strings.Replace(graph, "box", "rect", -1)
	page := fmt.Sprintf(template, graph)

	must(browser.OpenReader(bytes.NewBufferString(page)))
}
