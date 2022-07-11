package iterator

import (
	"testing"
)

func TestSliceIterator_HasNext(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				list: nil,
			},
			want: false,
		},
		{
			name: "0",
			fields: fields{
				list: make([]int, 0),
			},
			want: false,
		},
		{
			name: "n",
			fields: fields{
				list: []int{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSliceIterator[int](tt.fields.list)
			if got := s.HasNext(); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
