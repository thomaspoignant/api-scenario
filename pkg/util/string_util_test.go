package util_test

import (
	"testing"

	"github.com/thomaspoignant/api-scenario/pkg/util"
)

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty string", args{s: ""}, false},
		{"Int string", args{s: "1324"}, true},
		{"Float string", args{s: "1.1"}, true},
		{"Text string", args{s: "toto"}, false},
		{"Text string", args{s: "134toto"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
