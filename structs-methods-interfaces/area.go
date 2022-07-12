package main

type Rectangle struct {
	Width float64
	Height float64
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}