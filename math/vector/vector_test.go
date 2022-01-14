package vector

import (
	"testing"

	"github.com/puoklam/2dphysics/math/float"
)

func TestVectorAngle(t *testing.T) {
	tests := []struct {
		in   [2][]float64
		out  float64
		want bool
	}{
		{
			in:   [2][]float64{{1, 2}, {3, 4}},
			out:  0.17985349979247856,
			want: true,
		},
		{
			in:   [2][]float64{{5, 6, 7}, {8, 9, 10}},
			out:  0.04477778793657983,
			want: true,
		},
	}
	for i, tt := range tests {
		v1, v2 := NewVector(tt.in[0]...), NewVector(tt.in[1]...)
		got := float.Equal(Angle(v1, v2), tt.out)
		if got != tt.want {
			t.Errorf("%d. got %v; want %v", i, got, tt.want)
		}
	}
}

func BenchmarkCalAngle(b *testing.B) {
	vcts := [][2]*Vector{
		{NewVector(1, 2), NewVector(3, 4)},
		{NewVector(99.9, 128.99, 66.3), NewVector(1.2, 177.34, 10.6)},
		{NewVector(0.1, 0.2), NewVector(0.5, 0.6)},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Angle(vcts[i%len(vcts)][0], vcts[i%len(vcts)][1])
	}
	b.StopTimer()
}