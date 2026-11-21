package run

import (
	"testing"

	"encr.dev/pkg/option"
)

func TestTestTraceEndpoint(t *testing.T) {
	traced := option.Some("http://localhost:4000/trace")
	tests := []struct {
		name    string
		environ []string
		want    option.Option[string]
	}{
		{"unset", []string{"PATH=/bin"}, traced},
		{"disabled", []string{"ENCORE_TEST_DISABLE_TRACING=1"}, option.None[string]()},
		{"disabled true", []string{"ENCORE_TEST_DISABLE_TRACING=true"}, option.None[string]()},
		{"explicitly enabled", []string{"ENCORE_TEST_DISABLE_TRACING=0"}, traced},
		{"empty", []string{"ENCORE_TEST_DISABLE_TRACING="}, traced},
		{"invalid", []string{"ENCORE_TEST_DISABLE_TRACING=maybe"}, traced},
		{"last wins", []string{"ENCORE_TEST_DISABLE_TRACING=1", "ENCORE_TEST_DISABLE_TRACING=0"}, traced},
		{"prefix only", []string{"ENCORE_TEST_DISABLE_TRACING_X=1"}, traced},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testTraceEndpoint(tt.environ, 4000); got != tt.want {
				t.Errorf("testTraceEndpoint(%q) = %v, want %v", tt.environ, got, tt.want)
			}
		})
	}
}
