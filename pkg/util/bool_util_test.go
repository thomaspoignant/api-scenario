package util

import "testing"

func TestCompareBool(t *testing.T) {
	type args struct {
		expected bool
		value    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{ name: "Valid bool should return true",
			args: args{
				expected: true,
				value: "true",
			},
			want: true,
			wantErr: false,
		},
		{ name: "Valid bool should return false",
			args: args{
				expected: true,
				value: "false",
			},
			want: false,
			wantErr: false,
		},
		{ name: "Should return an error",
			args: args{
				expected: true,
				value: "fals",
			},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareBool(tt.args.expected, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}