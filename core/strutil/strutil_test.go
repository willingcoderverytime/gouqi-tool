package strutil

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		str string
	}
	var (
		tests = []struct {
			name string
			args args
			want bool
		}{
			{
				name: "this is  not empty",
				args: args{"sssss"},
				want: false,
			},
			{
				name: "this is   empty",
				args: args{""},
				want: true,
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(&tt.args.str); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

			






