package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	f := GetMessage()
	expected := string([]rune("Hello 🗺️ !"))
	if f != expected {
		t.Errorf("Unexpected result: Expected: %q Got: %q", expected, f)
	}

}
