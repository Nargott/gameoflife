package main

import (
	"reflect"
	"testing"
)

func TestNewPattern(t *testing.T) {
	type args struct {
		pattern [][]uint8
		width   int
		height  int
	}
	tests := []struct {
		name    string
		args    args
		want    *Pattern
		wantErr bool
	}{
		{
			name:    "Generate new pattern good case",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				width:   3,
				height:  3,
			},
			want:    &Pattern{
				p: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				w: 3,
				h: 3,
			},
			wantErr: false,
		},
		{
			name:    "Generate new pattern bad width",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				width:   0,
				height:  3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Generate new pattern bad height",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				width:   3,
				height:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Generate new pattern wrong height",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				width:   3,
				height:  4,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Generate new pattern wrong width",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				width:   2,
				height:  3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Generate new pattern wrong width at the row #2",
			args:    args{
				pattern: [][]uint8{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1},
				},
				width:   3,
				height:  3,
			},
			want:    nil,
			wantErr: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPattern(tt.args.pattern, tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPattern() got = %v, want %v", got, tt.want)
			}
		})
	}
}
