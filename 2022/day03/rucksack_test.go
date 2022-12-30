package day03

import (
	"io"
	"strings"
	"testing"
)

func TestReorganizeRucksack(t *testing.T) {
	type args struct {
		reader io.Reader
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
				reader: strings.NewReader(`
					vJrwpWtwJgWrhcsFMMfFFhFp
					jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
					PmmdzqPrVvPwwTWBwg
					wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
					ttgJtRGJQctTZtZT
					CrZsJsPPZsGzwwsLwLmpwMDw
				`),
			},
			want: 157,
		},
		{
			name: "Empty input",
			args: args{
				reader: strings.NewReader(``),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReorganizeRucksack(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReorganizeRucksack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReorganizeRucksack() = %v, want %v", got, tt.want)
			}
		})
	}
}
