package main

import "testing"

func Test_getNumberFromLine(t *testing.T) {
	type args struct {
		line string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{line: "1abc2"},
			want: 12,
		},
		{
			args: args{line: "pqr3stu8vwx"},
			want: 38,
		},
		{
			args: args{line: "a1b2c3d4e5f"},
			want: 15,
		},
		{
			args: args{line: "treb7uchet"},
			want: 77,
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
