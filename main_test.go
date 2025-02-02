package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Save the original stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	// Capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expectedOutput := "20\n0\n100\n1\n"
	if output != expectedOutput {
		t.Errorf("main() output = %v, want %v", output, expectedOutput)
	}
}

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
		{
			name: "test soma 3",
			args: args{
				a: 10,
				b: -10,
			},
			want: 0,
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

func TestSubtracao(t *testing.T) {
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
			name: "test subtracao 1",
			args: args{
				a: 10,
				b: 0,
			},
			want: 10,
		},
		{
			name: "test subtracao 2",
			args: args{
				a: 10,
				b: 10,
			},
			want: 0,
		},
		{
			name: "test subtracao 3",
			args: args{
				a: 10,
				b: -10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtracao(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplicacao(t *testing.T) {
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
			name: "test multiplicacao 1",
			args: args{
				a: 10,
				b: 0,
			},
			want: 0,
		},
		{
			name: "test multiplicacao 2",
			args: args{
				a: 10,
				b: 10,
			},
			want: 100,
		},
		{
			name: "test multiplicacao 3",
			args: args{
				a: 10,
				b: -10,
			},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplicacao(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiplicacao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivisao(t *testing.T) {
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
			name: "test divisao 1",
			args: args{
				a: 10,
				b: 0,
			},
			want: 0,
		},
		{
			name: "test divisao 2",
			args: args{
				a: 10,
				b: 10,
			},
			want: 1,
		},
		{
			name: "test divisao 3",
			args: args{
				a: 10,
				b: -10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divisao(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}
