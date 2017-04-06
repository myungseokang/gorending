package gorending

import (
	"testing"
)

func TestGorending(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gorending(); got != tt.want {
				t.Errorf("Gorending() = %v, want %v", got, tt.want)
			}
		})
	}
}
