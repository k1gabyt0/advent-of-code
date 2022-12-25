package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/k1gabyt0/AdventOfCode/2022/day01"
)

func TestSolvePuzzle(t *testing.T) {
	type args struct {
		in io.Reader
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example from AOC site",
			args: args{
				in: strings.NewReader(`
					1000
					2000
					3000

					4000

					5000
					6000

					7000
					8000
					9000

					10000
				`),
			},
			want: 24000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day01.CountMaxCalories(tt.args.in)
			if got != tt.want {
				t.Fatalf("got=%d cal, but wanted=%d cal", got, tt.want)
			}
		})
	}
}
