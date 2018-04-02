package staticModel

import (
	"testing"

	"github.com/ingmardrewing/staticPersistence"
)

func TestNewSiteDto(t *testing.T) {
	s := getSiteDto()
	actual := len(s.Containers())
	expected := 5

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestGetContainerByVariant(t *testing.T) {
	s := getSiteDto()

	actual := s.getContainerByVariant("marginal").Variant()
	expected := "marginal"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestGetPagesByVariant(t *testing.T) {
	s := getSiteDto()

	actual := len(s.getPagesByVariant("wurst", false))
	expected := 0

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	actual = len(s.getPagesByVariant("blog", false))
	expected = 23

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	actual = len(s.getPagesByVariant("blog", true))
	expected = 3

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPages(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Pages())
	expected := 0

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPosts(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Posts())
	expected := 23

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPostNaviPages(t *testing.T) {
	s := getSiteDto()

	actual := len(s.PostNaviPages())
	expected := 3

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestMarginals(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Marginals())
	expected := 3

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestNarratives(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Narratives())
	expected := 20

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestNarrativeMarginals(t *testing.T) {
	s := getSiteDto()

	actual := len(s.NarrativeMarginals())
	expected := 0

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestMarginalLocs(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Marginal())
	expected := 4

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestMainLocs(t *testing.T) {
	s := getSiteDto()

	actual := len(s.Main())
	expected := 3

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func getSiteDto() *siteDto {
	config := staticPersistence.ReadConfig("testResources", "configNew.json")
	return NewSiteDto(config[0]).(*siteDto)
}
