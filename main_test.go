package main

import "testing"

func Test_sum(t *testing.T){
	tc := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "ok",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "bad",
			a:    2,
			b:    4,
			want: 9,
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.a, tt.b)

			if got != tt.want {
				t.Error("test failed")
			}
		})
	}
}
