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

func TestConfigContainer(t *testing.T) {

	c := &configContainer{
		"twitterHandle",
		"topic",
		"tags",
		"cardType",
		"section",
		"fbPage",
		"twitterPage",
		"rssPath",
		"rssFilename",
		"css",
		"domain",
		"disqusId",
		"targetDir",
		"description",
		"homeText",
		"homeHeadline"}

	if c.TwitterHandle() != "twitterHandle" {
		t.Error("configCointainer returning wrong value on TwitterHandle()")
	}

	if c.Description() != "description" {
		t.Error("configCointainer returning wrong value on Desription()")
	}

	if c.Topic() != "topic" {
		t.Error("configCointainer returning wrong value on Topic()")
	}

	if c.Tags() != "tags" {
		t.Error("configCointainer returning wrong value on Tags()")
	}

	if c.Site() != "domain" {
		t.Error("configCointainer returning wrong value on Site()")
	}

	if c.CardType() != "cardType" {
		t.Error("configCointainer returning wrong value on CardType()")
	}

	if c.Section() != "section" {
		t.Error("configCointainer returning wrong value on Section()")
	}

	if c.FBPage() != "fbPage" {
		t.Error("configCointainer returning wrong value on FBPage()")
	}

	if c.TwitterPage() != "twitterPage" {
		t.Error("configCointainer returning wrong value on TwitterPage()")
	}

	if c.RssPath() != "rssPath" {
		t.Error("configCointainer returning wrong value on RssPath()")
	}

	if c.RssFilename() != "rssFilename" {
		t.Error("configCointainer returning wrong value on RssFilename()")
	}

	if c.Css() != "css" {
		t.Error("configCointainer returning wrong value on Css()")
	}

	if c.DisqusId() != "disqusId" {
		t.Error("configCointainer returning wrong value on DisqusId()")
	}

	if c.TargetDir() != "targetDir" {
		t.Error("configCointainer returning wrong value on TargetDir()")
	}
}
