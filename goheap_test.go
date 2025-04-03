package goheap_test

import (
	"fmt"
	"testing"

	"github.com/uulwake/goheap"
)

func TestMinHeap(t *testing.T) {
	minHeap := goheap.NewHeap(func(a, b int) bool { return a < b }, -1)

	inputs := []int{10, 5, 8, 3, 1, 2, 6}
	for _, input := range inputs {
		minHeap.Enqueue(input)
	}

	tests := []struct {
		name    string
		value   int
		size    int
		isEmpty bool
	}{
		{"first", 1, 6, false},
		{"second", 2, 5, false},
		{"third", 3, 4, false},
		{"fourth", 5, 3, false},
		{"fifth", 6, 2, false},
		{"sixth", 8, 1, false},
		{"seventh", 10, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peek, ok := minHeap.Peek()

			if !ok {
				t.Errorf("expect peek value to be ok: %d", peek)
			}

			if peek != tt.value {
				t.Errorf("expect peek %d but got %d", tt.value, peek)
			}

			el, ok := minHeap.Dequeue()

			if !ok {
				t.Errorf("expect dequeue value to be ok: %d", el)
			}

			if el != tt.value {
				t.Errorf("expect dequeue %d but got %d", tt.value, el)
			}

			if minHeap.Size() != tt.size {
				t.Errorf("expect size %d but got %d", tt.size, minHeap.Size())
			}

			if minHeap.IsEmpty() != tt.isEmpty {
				t.Errorf("expect isEmpty %t but got %t", tt.isEmpty, minHeap.IsEmpty())
			}
		})
	}
}

func TestMaxHeap(t *testing.T) {
	minHeap := goheap.NewHeap(func(a, b int) bool { return a > b }, -1)

	inputs := []int{1, 5, 8, 3, 10, 2, 6}
	for _, input := range inputs {
		minHeap.Enqueue(input)
	}

	tests := []struct {
		name    string
		value   int
		size    int
		isEmpty bool
	}{
		{"seventh", 10, 6, false},
		{"sixth", 8, 5, false},
		{"fifth", 6, 4, false},
		{"fourth", 5, 3, false},
		{"third", 3, 2, false},
		{"second", 2, 1, false},
		{"first", 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peek, ok := minHeap.Peek()

			if !ok {
				t.Errorf("expect peek value to be ok: %d", peek)
			}

			if peek != tt.value {
				t.Errorf("expect peek %d but got %d", tt.value, peek)
			}

			el, ok := minHeap.Dequeue()

			if !ok {
				t.Errorf("expect dequeue value to be ok: %d", el)
			}

			if el != tt.value {
				t.Errorf("expect dequeue %d but got %d", tt.value, el)
			}

			if minHeap.Size() != tt.size {
				t.Errorf("expect size %d but got %d", tt.size, minHeap.Size())
			}

			if minHeap.IsEmpty() != tt.isEmpty {
				t.Errorf("expect isEmpty %t but got %t", tt.isEmpty, minHeap.IsEmpty())
			}
		})
	}
}

type Person struct {
	name   string
	weight int
}

func TestMinHeapStruct(t *testing.T) {
	minHeap := goheap.NewHeap(func(a, b Person) bool { return a.weight < b.weight }, Person{})

	weights := []int{10, 5, 8, 3, 1, 2, 6}
	for i, weight := range weights {
		minHeap.Enqueue(Person{name: fmt.Sprintf("Person %d", i+1), weight: weight})
	}

	tests := []struct {
		name    string
		value   int
		size    int
		isEmpty bool
	}{
		{"first", 1, 6, false},
		{"second", 2, 5, false},
		{"third", 3, 4, false},
		{"fourth", 5, 3, false},
		{"fifth", 6, 2, false},
		{"sixth", 8, 1, false},
		{"seventh", 10, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peek, ok := minHeap.Peek()

			if !ok {
				t.Errorf("expect peek value to be ok: %d", peek.weight)
			}

			if peek.weight != tt.value {
				t.Errorf("expect peek %d but got %d", tt.value, peek.weight)
			}

			el, ok := minHeap.Dequeue()

			if !ok {
				t.Errorf("expect dequeue value to be ok: %d", el.weight)
			}

			if el.weight != tt.value {
				t.Errorf("expect dequeue %d but got %d", tt.value, el.weight)
			}

			if minHeap.Size() != tt.size {
				t.Errorf("expect size %d but got %d", tt.size, minHeap.Size())
			}

			if minHeap.IsEmpty() != tt.isEmpty {
				t.Errorf("expect isEmpty %t but got %t", tt.isEmpty, minHeap.IsEmpty())
			}
		})
	}
}

func TestMaxHeapStruct(t *testing.T) {
	minHeap := goheap.NewHeap(func(a, b Person) bool { return a.weight > b.weight }, Person{})

	weights := []int{1, 5, 8, 3, 10, 2, 6}
	for i, weight := range weights {
		minHeap.Enqueue(Person{name: fmt.Sprintf("Person %d", i+1), weight: weight})
	}

	tests := []struct {
		name    string
		value   int
		size    int
		isEmpty bool
	}{
		{"seventh", 10, 6, false},
		{"sixth", 8, 5, false},
		{"fifth", 6, 4, false},
		{"fourth", 5, 3, false},
		{"third", 3, 2, false},
		{"second", 2, 1, false},
		{"first", 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peek, ok := minHeap.Peek()

			if !ok {
				t.Errorf("expect peek value to be ok: %d", peek.weight)
			}

			if peek.weight != tt.value {
				t.Errorf("expect peek %d but got %d", tt.value, peek.weight)
			}

			el, ok := minHeap.Dequeue()

			if !ok {
				t.Errorf("expect dequeue value to be ok: %d", el.weight)
			}

			if el.weight != tt.value {
				t.Errorf("expect dequeue %d but got %d", tt.value, el.weight)
			}

			if minHeap.Size() != tt.size {
				t.Errorf("expect size %d but got %d", tt.size, minHeap.Size())
			}

			if minHeap.IsEmpty() != tt.isEmpty {
				t.Errorf("expect isEmpty %t but got %t", tt.isEmpty, minHeap.IsEmpty())
			}
		})
	}
}
