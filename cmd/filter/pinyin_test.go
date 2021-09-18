package main

import (
	"testing"
)

func Test_toPinyinKey(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "左半屏", args: args{"左半屏"}, want: "zbp"},
		{name: "左侧2/3", args: args{"左侧2/3"}, want: "zc2/3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPinyinKey(tt.args.s); got != tt.want {
				t.Errorf("toPinyinKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
