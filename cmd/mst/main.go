package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/msiebuhr/mst"
)

func main() {
	var wg sync.WaitGroup
	flag.Parse()

	// Start go-routine to process numbers
	ch := make(chan float64, 100)
	data := mst.NewData()
	wg.Add(1)
	go data.AddChan(ch, &wg)

	// Read lines, parse numbers and do maths
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

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

	names := []string{
		"min",
		"q1",
		"median",
		"q3",
		"max",
		"average",
		"sum",
		"count",
		"stddev",
	}

	// Write output data
	stats, err := data.GetStatistics(names)

	if err != nil {
		panic(err)
	}

	for _, name := range names {
		fmt.Println(name, stats[name])
	}
}
