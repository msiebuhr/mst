package main

import (
	"math"
	"sync"
	"sort"
)

type Data struct {
	max   float64
	min   float64
	sum   float64
	count float64
	data  []float64
}

func NewData() Data {
	return Data{
		max:   -1 * math.MaxFloat64,
		min:   math.MaxFloat64,
		sum:   0,
		count: 0,
		data:  make([]float64, 0),
	}
}

// Add a number, but don't do array-sorting & sunc
func (d *Data) AddNumber(n float64) {
	d.data = append(d.data, n)

	if n > d.max {
		d.max = n
	}
	if n < d.min {
		d.min = n
	}
	d.count += 1
	d.sum += n
}

// Do some optimization on data before marching on
func (d *Data) Finalize() {
	// Sort data - later computations expect it..
	sort.Float64s(d.data)
}

func (d *Data) in(in <-chan float64, wg *sync.WaitGroup) {
	defer d.Finalize()
	defer wg.Done()

	for n := range in {
		d.AddNumber(n)
	}
}

func (d *Data) Average() float64 {
	return d.sum / d.count
}

func (d *Data) Percentile(which float64) float64 {
	index := int(math.Trunc(float64(len(d.data)) * which))
	return d.data[index]
}