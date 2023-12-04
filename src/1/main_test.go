package main

import "testing"

func Test_getNumberFromLine(t *testing.T) {
	type args struct {
		line string
	}

	tests := []struct {
		name string
		args args
		want uint16
	}{
		// TODO: Add test cases.
		{
			name: "Test with number in line",
			args: args{line: "The number is 123."},
			want: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberFromLine(tt.args.line); got != tt.want {
				t.Errorf("getNumberFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
