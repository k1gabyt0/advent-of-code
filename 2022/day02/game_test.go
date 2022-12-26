package day02_test

import (
	"io"
	"strings"
	"testing"

	"github.com/k1gabyt0/AdventOfCode/2022/day02"
)

func TestCalculateTotalScore(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Empty input",
			args: args{
				in: strings.NewReader(""),
			},
			want: 0,
		},
		{
			name: "Example from AOC site",
			args: args{
				in: strings.NewReader(`
					A Y
					B X
					C Z
				`),
			},
			want: 15,
		},
		{
			name: "Example from AOC site without new linen at start and at end",
			args: args{
				in: strings.NewReader(`A Y
					B X
					C Z`),
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day02.CalculateTotalScore(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTotalScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTotalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
