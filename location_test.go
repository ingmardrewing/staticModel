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

func TestDomain(t *testing.T) {
	l := getLocation()

	l.Domain("testDomain")

	if l.Domain() != "testDomain" {
		t.Error("Expected loc to have the newly set value in domain field.")
	}
}

func TestHtmlFilename(t *testing.T) {
	l := getLocation()

	l.HtmlFilename("testHtmlFilename")

	if l.HtmlFilename() != "testHtmlFilename" {
		t.Error("Expected loc to have the newly set value in htmlFilename field.")
	}
}

func TestTitle(t *testing.T) {
	l := getLocation()

	l.Title("testTitle")

	if l.Title() != "testTitle" {
		t.Error("Expected loc to have the newly set value in title.")
	}
}

func TestThumbnailUrl(t *testing.T) {
	l := getLocation()

	l.ThumbnailUrl("testUrl")

	if l.ThumbnailUrl() != "testUrl" {
		t.Error("Expected loc to have the newly set value in thumbnailUrl.")
	}
}

func TestExternalLinks(t *testing.T) {
	l := getLocation()

	l.ExternalLink("externalLink")

	if l.ExternalLink() != "externalLink" {
		t.Error("Expected loc to have the newly set value in externalLink.")
	}
}

func getLocation() *loc {
	return NewLocation("https://example.com", "example.com", "test-title", "thumbnailUrl", "pathFromDocRoot", "filename")
}
