package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Pop(t *testing.T) {
	type args struct {
		n []byte
	}
	type want struct {
		v   byte
		ok  bool
		len int
		cap int
	}

	tests := []struct {
		name string
		s    *Stack
		args args
		want want
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			s:    NewStack(0),
			args: args{},
			want: want{' ', false, 0, 0},
		},
		{
			name: "5 elements",
			s:    NewStack(8),
			args: args{[]byte{0, 1, 2, 3, 4}},
			want: want{4, true, 4, 8},
		},
		{
			name: "10 elements",
			s:    NewStack(8),
			args: args{[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: want{9, true, 9, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.n...)
			v, ok := tt.s.Pop()
			w := want{v, ok, tt.s.Len(), tt.s.Cap()}
			if tt.want != w {
				t.Errorf("Stack.Pop() got = %v, want %v", w, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		n []byte
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want int
	}{

		{
			name: "0 cap",
			s:    NewStack(0),
			args: args{[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: 16,
		},
		{
			name: "8 cap",
			s:    NewStack(8),
			args: args{[]byte{0, 1, 2, 3, 4, 5}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.n...)
			assert.Equal(t, tt.want, tt.s.Cap())
		})
	}
}
