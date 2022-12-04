package daythree

import "testing"

func Test_itemType_getPriority(t *testing.T) {
	tests := []struct {
		name string
		it   itemType
		want int
	}{
		{
			name: "a",
			it: "a",
			want: 1,
		},
		{
			name: "z",
			it: "z",
			want: 26,
		},
		{
			name: "A",
			it: "A",
			want: 27,
		},
		{
			name: "Z",
			it: "Z",
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.it.getPriority(); got != tt.want {
				t.Errorf("itemType.getPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
