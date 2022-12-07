package dayfive

import (
	"reflect"
	"testing"
)

func Test_prepend(t *testing.T) {
	type args struct {
		s         []string
		container string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one item",
			args: args{
				s:         []string{"B", "C", "D"},
				container: "A",
			},
			want: []string{"A", "B", "C", "D"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepend(tt.args.s, tt.args.container); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{
			name: "happy path",
			args: args{
				s: "move 3 from 4 to 3",
			},
			want: instruction{
				fromStack: 4,
				toStack: 3,
				numMoves: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
