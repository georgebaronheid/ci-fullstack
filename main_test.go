package main

import "testing"

func TestSoma(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test soma 1",
			args: args{
				a: 10,
				b: 0,
			},
			want: 10,
		},
		{
			name: "test soma 2",
			args: args{
				a: 10,
				b: 10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Soma(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Soma() = %v, want %v", got, tt.want)
			}
		})
	}
}
