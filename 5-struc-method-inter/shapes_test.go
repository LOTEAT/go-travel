package main

import (
	"testing"
)

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func Perimeter2(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func Area2(rect Rectangle) float64 {
	return rect.width * rect.height
}

func TestPerimeter2(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter2(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
