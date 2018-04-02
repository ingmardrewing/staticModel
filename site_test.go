package staticModel

import (
	"testing"

	"github.com/ingmardrewing/staticPersistence"
)

func TestNewSiteDto(t *testing.T) {
	config := staticPersistence.ReadConfig("testResources", "configNew.json")
	siteDto := NewSiteDto(config[0])

	actual := len(siteDto.Containers())
	expected := 5

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
