package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Benchmark struct {
	tests map[string]int
	name  string
}

func NewBenchmark(name string) *Benchmark {
	return &Benchmark{
		tests: make(map[string]int),
		name:  name,
	}
}

func (b *Benchmark) AddTest(name string, result int) {
	b.tests[name] = result
}

func ParseTest(test string) (name string, value int, valid bool) {
	lowerName := strings.ToLower(test)
	if !strings.HasPrefix(lowerName, "benchmark") {
		return "", 0, false
	}

	tokens := strings.Fields(lowerName)

	if tokens == nil {
		return "", 0, false
	}

	if len(tokens) != 4 {
		return "", 0, false
	}

	if tokens[3] != "ns/op" {
		return "", 0, false
	}

	value, err := strconv.Atoi(tokens[2])
	if err != nil {
		return "", 0, false
	}

	return tokens[0], value, true
}

func LoadBenchmarkResults(filename string) (*Benchmark, error) {

	benchmark := NewBenchmark(filename)

	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		name, value, valid := ParseTest(scanner.Text())
		if valid {
			benchmark.AddTest(name, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return benchmark, nil
}

type CompareResult struct {
	name  string
	value float32
}

func NewCompareResult(name string, result float32) *CompareResult {
	return &CompareResult{
		name:  name,
		value: result,
	}
}

type Comparison struct {
	result []*CompareResult
}

func NewComparison() *Comparison {
	return &Comparison{
		result: make([]*CompareResult, 0),
	}
}

func (c *Comparison) AddResult(name string, result float32) {
	c.result = append(c.result, NewCompareResult(name, result))
}

func (c *Comparison) String() string {

	buffer := new(bytes.Buffer)
	w := tabwriter.NewWriter(buffer, 0, 8, 0, '\t', tabwriter.AlignRight)

	fmt.Fprintln(w, "Improvment\tBenchmark")

	sum := float32(0)
	count := len(c.result)

	for _, result := range c.result {
		fmt.Fprintf(w, "%.2f %%\t%s\n", result.value, result.name)
		sum += result.value
	}

	fmt.Fprintln(w, "\t")

	if count > 0 {
		fmt.Fprintf(w, "%.2f %%\tTotal improvement\n", sum/float32(count))
	}

	w.Flush()

	return buffer.String()
}

func CompareBenchmarks(old *Benchmark, current *Benchmark) *Comparison {
	if old == nil || current == nil {
		return nil
	}

	comparison := NewComparison()

	for name, value := range current.tests {
		if old_val := old.tests[name]; old_val != 0 {
			improvment := float32(old_val) / float32(value) * float32(100)
			comparison.AddResult(name, improvment)
		}
	}

	return comparison
}
