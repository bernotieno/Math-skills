package calculation

import "testing"

func TestStandDev(t *testing.T) {
	type args struct {
		deveation float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for an odd length of slice",
			args: args{
				deveation: Variance([]float64{12, 43, 54, 29, 95, 84, 74}),
			},
			want: 27.946012091830465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandDev(tt.args.deveation); got != tt.want {
				t.Errorf("StandDev() = %v, want %v", got, tt.want)
			}
		})
	}
}
