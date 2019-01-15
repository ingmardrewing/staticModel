package staticModel

import "testing"

func TestNewLocation(t *testing.T) {
	l := getLocation()

	expected := "https://example.com"
	actual := l.extLink

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPathFromDocRootWithName(t *testing.T) {
	l := getLocation()

	expected := "/pathFromDocRoot/filename"
	actual := l.PathFromDocRootWithName()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestUrl(t *testing.T) {
	l := getLocation()

	expected := "https://example.com/pathFromDocRoot/filename"
	actual := l.Url()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func getLocation() *loc {
	return NewLocation("https://example.com", "example.com", "test-title", "thumbnailUrl", "pathFromDocRoot", "filename", "microThumbnailUrl")
}
