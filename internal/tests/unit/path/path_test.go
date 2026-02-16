package path

import (
	"testing"

	"github.com/pashkov256/deletor/internal/path"
)

func TestPathVariables(t *testing.T) {
	tests := []struct {
		name      string
		got, want any
	}{
		{"AppDirName", path.AppDirName, "deletor"},
		{"RuleFileName", path.RuleFileName, "rule.json"},
		{"LogFileName", path.LogFileName, "deletor.log"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
