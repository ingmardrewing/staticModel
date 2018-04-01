package staticModel

import "testing"

func TestNewLocation(t *testing.T) {
	l := NewLocation("https://example.com", "example.com", "test-title", "thumbnailUrl", "pathFromDocRoot", "filename")

	expected := "https://example.com"
	actual := l.extLink

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
