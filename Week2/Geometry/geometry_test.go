package geometry

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	check := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{4.0, 3.0}
		want := 12.0
		check(t, rect, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		check(t, circle, 314.1592653589793)
	})

}

func TestDrivenTest(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, val := range areaTests {
		got := val.shape.Area()
		if got != val.want {
			t.Errorf("got %v want %v", got, val.want)
		}
	}

}
