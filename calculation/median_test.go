package calculation

import "testing"

func TestMedian(t *testing.T) {
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
			args: args{numbers: []float64{12, 43, 54, 29, 95, 84, 74}},
			want: 54,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Median(tt.args.numbers)
			if got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
