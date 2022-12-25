package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/k1gabyt0/AdventOfCode/2022/day01"
)

func TestSolvePuzzle(t *testing.T) {
	type args struct {
		in           io.Reader
		leadersCount uint
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
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
				leadersCount: 1,
			},
			want: 24000,
		},
		{
			name: "Empty record",
			args: args{
				in:           strings.NewReader(""),
				leadersCount: 1,
			},
			want: 0,
		},
		{
			name: "Only one record",
			args: args{
				in:           strings.NewReader("1000"),
				leadersCount: 1,
			},
			want: 1000,
		},
		{
			name: "0 leaders count should return error",
			args: args{
				in:           strings.NewReader("1000"),
				leadersCount: 0,
			},
			wantErr: true,
		},
		{
			name: "Top 3 elves",
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
				leadersCount: 3,
			},
			want: 45000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day01.CountMaxCalories(tt.args.in, tt.args.leadersCount)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, but didn't get it")
				}
				if got != 0 {
					t.Fatalf("expected zero value, but got=%d", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error=%q", err)
			}

			if got != tt.want {
				t.Fatalf("got=%d cal, but wanted=%d cal", got, tt.want)
			}
		})
	}
}
