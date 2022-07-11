package iterator

import (
	"container/list"
	"testing"
)

func TestListIterator_HasNext(t *testing.T) {
	type fields struct {
		list *list.List
	}
	caseN := list.New()
	caseN.PushBack(1)
	caseN.PushBack(2)
	caseN.PushBack(3)
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
				list: list.New(),
			},
			want: false,
		},
		{
			name: "n",
			fields: fields{
				list: caseN,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewListIterator[int](tt.fields.list)
			if got := s.HasNext(); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
