package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	flag.Parse()

	// Start go-routine to process numbers
	ch := make(chan float64, 100)
	data := NewData()
	wg.Add(1)
	go data.in(ch, &wg)

	// Read lines, parse numbers and do maths
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error parsing number:", err)
		}

		ch <- value
	}
	close(ch)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	wg.Wait()

	// Write output data
	fmt.Println("min", data.min)
	fmt.Println("q1", data.Percentile(0.25))
	fmt.Println("mean", data.Percentile(0.5))
	fmt.Println("q3", data.Percentile(0.75))
	fmt.Println("max", data.max)
	fmt.Println("sum", data.sum)
	fmt.Println("count", data.count)
	fmt.Println("average", data.sum/data.count)
}
