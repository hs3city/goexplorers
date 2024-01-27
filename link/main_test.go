package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_catch(t *testing.T) {
	type args struct {
		val []byte
		err error
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catch(tt.args.val, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catch() = %v, want %v", got, tt.want)
			}
		})
	}
}
