package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
        name string
		shape Shape
		want  float64
	}{
        {name: "rectangle", shape: Rectangle{10.0, 10.0}, want: 100},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{4, 2}, want: 4},
	}

	for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()

            if got != tt.want {
                t.Errorf("%#v got %g, want %g", tt.shape, got, tt.want)
            }
        })
	}
}
