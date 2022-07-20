package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mcjcloud/wool/graph"
)

var src = flag.String("src", "", "starting dependency")
var dst = flag.String("dst", "", "dependency to connect to")

func main() {
	flag.Parse()

	graphOut, err := readPipe()
	if err != nil {
		log.Fatalf("error reading from pipe: %s", err)
	}

	fmt.Print("constructing graph... ")
	g, err := graph.New(graphOut)
	if err != nil {
		log.Fatalf("error constructing graph: %s", err)
	}
	fmt.Println("done.")

	paths, err := g.Trace(*src, *dst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("found %d paths.\n", len(paths))
	for _, path := range paths {
		fmt.Printf("> %d hops\n", len(path))
		fmt.Print(path[0].Value)
		for _, p := range path[1:] {
			fmt.Printf(" -> %s", p.Value)
		}
		fmt.Println()
	}
}

func readPipe() (res string, err error) {
	if !isInputFromPipe() {
		err = fmt.Errorf("expected piped output from 'go mod graph'")
		return
	}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		if res != "" {
			res += "\n"
		}
		res += scanner.Text()
	}
	return
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
