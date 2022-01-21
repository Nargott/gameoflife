package main

import (
	"math"
	"reflect"
	"testing"
)

func TestNewLife(t *testing.T) {
	pattern, _ := NewPattern([][]uint8{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}, 3, 3)

	const (
		width = 25
		height = 25
	)
	c := newUniverse(width, height)

	c.spawnPattern(pattern, int(math.Round(float64(25)/2)), int(math.Round(float64(25)/2)))

	type args struct {
		width   int
		height  int
		pattern *Pattern
	}
	tests := []struct {
		name    string
		args    args
		want    *Life
		wantErr bool
	}{
		{
			name:    "Create new life good case",
			args:    args{
				width:   width,
				height:  height,
				pattern: pattern,
			},
			want:    &Life{
				current: c,
				next:    newUniverse(width, height),
				w:       width,
				h:       height,
			},
			wantErr: false,
		},
		{
			name:    "Create new life no pattern",
			args:    args{
				width:   width,
				height:  height,
				pattern: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Create new life wrong width",
			args:    args{
				width:   1,
				height:  height,
				pattern: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Create new life wrong height",
			args:    args{
				width:   width,
				height:  1,
				pattern: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLife(tt.args.width, tt.args.height, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLife() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLife() got = %v, want %v", got, tt.want)
			}
		})
	}
}
