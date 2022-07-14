package main

import "testing"


func TestArea(t *testing.T) {

	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
	}

	for _, tt := range areaTest {

		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
// func TestArea(t *testing.T) {

// 	checkArea := func (t testing.TB, shape Shape, want float64)  {

// 		t.Helper()

// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
	
// 	t.Run("rectangle", func(t *testing.T) {
		
// 		rectangle := Rectangle{12.0, 6.0}

// 		checkArea(t, rectangle, 72.0)
// 	});
	
// 	t.Run("circle", func(t *testing.T) {

// 		circle := Circle{10.0}

// 		checkArea(t, circle, 314.1592653589793)
// 	});
// }
// func TestArea(t *testing.T) {

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}

// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circle", func(t *testing.T) {

// 		circle := Circle{10}

// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {
// 	t.Run("rectangle", func(t *testing.T) {

// 		rectangle := Rectangle{12.0, 6.0}

// 		got := Area(rectangle)
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}

// 		got := Area(circle)
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }
