package dayfour

import "testing"

func Test_elfCleaningAssignmentPairs_hasFullyContainedAssignment(t *testing.T) {
	type fields struct {
		firstElf  sectionList
		secondElf sectionList
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "first contains second",
			fields: fields{
				firstElf: sequentialInt(6,10),
				secondElf: sequentialInt(6,6),
			},
			want: true,
		},
		{
			name: "second contains first",
			fields: fields{
				firstElf: sequentialInt(6,6),
				secondElf: sequentialInt(6,10),
			},
			want: true,
		},
		{
			name: "first contains most of second",
			fields: fields{
				firstElf: sequentialInt(6,10),
				secondElf: sequentialInt(7,11),
			},
			want: true,
		},
		{
			name: "first contains one of second",
			fields: fields{
				firstElf: sequentialInt(6,10),
				secondElf: sequentialInt(10,11),
			},
			want: true,
		},
		{
			name: "unrelated",
			fields: fields{
				firstElf: sequentialInt(88,100),
				secondElf: sequentialInt(7,11),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &elfCleaningAssignmentPairs{
				firstElf:  tt.fields.firstElf,
				secondElf: tt.fields.secondElf,
			}
			if got := p.hasFullyContainedAssignment(); got != tt.want {
				t.Errorf("elfCleaningAssignmentPairs.hasFullyContainedAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
