package main

import ("testing"
"sort"
)

func TestSimple(t *testing.T) {
	d := NewData()

	d.AddNumber(2)
	d.AddNumber(4)
	d.AddNumber(3)
	d.AddNumber(5)
	d.AddNumber(1)
	d.Finalize()

if !sort.Float64sAreSorted(d.data) {
	t.Errorf("Expected data to be sorted after finalization - they are not: %v", d.data)
}

	if d.min != 1 {
		t.Errorf("Expected min to be 1, got %f", d.min)
	}

	if d.max != 5 {
		t.Errorf("Expected max to be 5, got %f", d.max)
	}

	if d.sum != 15 {
		t.Errorf("Expected sum to be 15, got %f", d.sum)
	}

	if d.Percentile(0.5) != 3 {
		t.Errorf("Expected mean (0.5 percentile) to be 3, got %f", d.Percentile(0.5))
	}

	if d.Percentile(0.25) != 2 {
		t.Errorf("Expected q1 (0.25 percentile) to be 2, got %f", d.Percentile(0.25))
	}

	if d.Percentile(0.75) != 4 {
		t.Errorf("Expected q4 (0.75 percentile) to be 4, got %f", d.Percentile(0.75))
	}
}
