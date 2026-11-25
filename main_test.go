package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "обычный размер", size: 100},
		{name: "один элемент", size: 1},
		{name: "нулевой размер", size: 0},
		{name: "отрицательный размер", size: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			expectedLength := tt.size
			if expectedLength < 0 {
				expectedLength = 0
			}

			if len(data) != expectedLength {
				t.Fatalf("получена длина %d, ожидалась %d", len(data), expectedLength)
			}

			for _, value := range data {
				if value <= 0 || value > SIZE {
					t.Errorf("число %d вне диапазона от 1 до %d", value, SIZE)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 3, 7, 1}, want: 9},
		{name: "максимум в конце", data: []int{1, 3, 7, 9}, want: 9},
		{name: "одинаковые максимумы", data: []int{4, 10, 2, 10}, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{17}, want: 17},
		{name: "меньше восьми элементов", data: []int{2, 8, 1}, want: 8},
		{name: "ровно восемь элементов", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "длина не делится на восемь", data: []int{4, 7, 2, 11, 3, 6, 1, 9, 5, 10}, want: 11},
		{name: "максимум в последней части", data: []int{1, 2, 3, 4, 5, 6, 7, 100}, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}
