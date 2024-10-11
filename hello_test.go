package gotpl

import "testing"

func TestSay(t *testing.T) {
	expected := "Hello, World!"
	result := Say()
	if result != expected {
		t.Errorf("Say() = %v; want %v", result, expected)
	}
}
