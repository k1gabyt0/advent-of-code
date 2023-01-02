package day03

import (
	"io"
	"strings"
	"testing"
)

func TestReorganizeRucksack(t *testing.T) {
	type args struct {
		reader    io.Reader
		groupSize int
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
				groupSize: 3,
			},
			want: 70,
		},
		{
			name: "Empty input",
			args: args{
				reader:    strings.NewReader(``),
				groupSize: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReorganizeRucksack(tt.args.reader, tt.args.groupSize)
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
