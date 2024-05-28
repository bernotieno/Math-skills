package calculation

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for an odd length of slice",
			args: args{
				numbers: []float64{12, 43, 54, 29, 95, 84, 74},
			},
			want: 780.9795918367346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.numbers); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
