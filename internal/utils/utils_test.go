package utils

import (
	"reflect"
	"testing"
)

func TestParseExtToSlice(t *testing.T) {
	tests := []struct {
		name       string
		extensions string
		want       []string
	}{
		{
			name:       "Basic valid extensions",
			extensions: "jpg,png,gif",
			want:       []string{".jpg", ".png", ".gif"},
		},
		{
			name:       "Extensions with existing dot prefixes",
			extensions: ".jpg,.png,.gif",
			want:       []string{".jpg", ".png", ".gif"},
		},
		{
			name:       "Mixed casing and extra whitespace",
			extensions: " JPG , .Png , Gif ",
			want:       []string{".jpg", ".png", ".gif"},
		},
		{
			name:       "Empty segments",
			extensions: "jpg,,png",
			want:       []string{".jpg", ".png"},
		},
		{
			name:       "Empty string",
			extensions: "",
			want:       []string{},
		},
		{
			name:       "Whitespace-only string",
			extensions: "   ",
			want:       []string{},
		},
		{
			name:       "Multiple dots and special chars",
			extensions: "..jpg, .tar.gz, txt ",
			want:       []string{"..jpg", ".tar.gz", ".txt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseExtToSlice(tt.extensions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseExtToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseExcludeToSlice(t *testing.T) {
	tests := []struct {
		name    string
		exclude string
		want    []string
	}{
		{
			name:    "Basic valid patterns",
			exclude: "node_modules,vendor,temp",
			want:    []string{"node_modules", "vendor", "temp"},
		},
		{
			name:    "Mixed casing and extra whitespace",
			exclude: " node_modules , Vendor , TEMP ",
			want:    []string{"node_modules", "Vendor", "TEMP"},
		},
		{
			name:    "Empty segments",
			exclude: "a,,b",
			want:    []string{"a", "b"},
		},
		{
			name:    "Empty string",
			exclude: "",
			want:    []string{},
		},
		{
			name:    "Whitespace-only string",
			exclude: "   ",
			want:    []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseExcludeToSlice(tt.exclude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseExcludeToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
