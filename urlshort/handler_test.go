package urlshort

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMapHandler(t *testing.T) {
	type args struct {
		pathsToUrls map[string]string
		fallback    http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapHandler(tt.args.pathsToUrls, tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYAMLHandler(t *testing.T) {
	type args struct {
		yml      []byte
		fallback http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    http.HandlerFunc
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLHandler(tt.args.yml, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAMLHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONHandler(t *testing.T) {
	type args struct {
		jsonInput []byte
		fallback  http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    http.HandlerFunc
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONHandler(tt.args.jsonInput, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
