package main

import (
	"github.com/binaryphile/conway/tbox"
	"reflect"
	"testing"
)

func Test_generateOutput(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateOutput(); got != tt.want {
				t.Errorf("generateOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gridFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tbox.gridFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initialize(t *testing.T) {
	tests := []struct {
		name    string
		want    func()
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTermbox()
			if (err != nil) != tt.wantErr {
				t.Errorf("initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_show(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbox.Show(tt.args.output)
		})
	}
}

func Test_showGrid(t *testing.T) {
	type args struct {
		grid [][]rune
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbox.showGrid(tt.args.grid)
		})
	}
}

func Test_waitForInput(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbox.WaitForInput()
		})
	}
}
