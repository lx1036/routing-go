package main

import "testing"

func Perimeter(width float64, height float64) float64 {
	return (width + height) * 2
}

func Area(width float64, height float64) float64 {
	return width * height
}

type Rectangle struct {
	width float64
	height float64
}

func TestPerimeter(t *testing.T)  {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // .2 表示输出两位小数, f 表示 float64
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
