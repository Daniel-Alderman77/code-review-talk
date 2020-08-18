package geometry

import "testing"

func TestTriangle(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Equilateral", args{3, 3, 3}, "Equilateral"},
		{"Isosceles", args{5, 5, 4}, "Isosceles"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Triangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Triangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
