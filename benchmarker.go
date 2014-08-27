package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: benchmarker old_resuls.log new_results.log\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	filenames := flag.Args()
	if len(filenames) != 2 {
		flag.Usage()
		log.Fatal("Invalid input parameters.")
	}

	prev_bench, err := LoadBenchmarkResults(filenames[0])
	if err != nil {
		log.Fatal(err)
	}

	new_bench, err := LoadBenchmarkResults(filenames[1])
	if err != nil {
		log.Fatal(err)
	}

	results := CompareBenchmarks(prev_bench, new_bench)

	sort.Sort(results)
	fmt.Println(results)
}
