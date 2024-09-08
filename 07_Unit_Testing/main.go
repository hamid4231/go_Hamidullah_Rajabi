package main

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length   int
		width    int
		expected string
	}{
		{length: 4, width: 2, expected: "even rectangle"},
		{length: 5, width: 3, expected: "odd rectangle"},
		{length: 6, width: 4, expected: "even rectangle"},
		{length: 7, width: 1, expected: "odd rectangle"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := RectangleArea(tt.length, tt.width)
			if result != tt.expected {
				t.Errorf("RectangleArea(%d, %d) = %v; want %v", tt.length, tt.width, result, tt.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length   int
		width    int
		expected bool
	}{
		{length: 4, width: 6, expected: true},
		{length: 5, width: 2, expected: true},
		{length: 3, width: 3, expected: false},
		{length: 7, width: 8, expected: true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := RectanglePerimeter(tt.length, tt.width)
			if result != tt.expected {
				t.Errorf("RectanglePerimeter(%d, %d) = %v; want %v", tt.length, tt.width, result, tt.expected)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		s        int
		expected string
	}{
		{s: 4, expected: "even square"},
		{s: 3, expected: "odd square"},
		{s: 2, expected: "even square"},
		{s: 7, expected: "odd square"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := SquareArea(tt.s)
			if result != tt.expected {
				t.Errorf("SquareArea(%d) = %v; want %v", tt.s, result, tt.expected)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		s        int
		expected bool
	}{
		{s: 10, expected: true},
		{s: 9, expected: true},
		{s: 5, expected: false},
		{s: 2, expected: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := SquarePerimeter(tt.s)
			if result != tt.expected {
				t.Errorf("SquarePerimeter(%d) = %v; want %v", tt.s, result, tt.expected)
			}
		})
	}
}
