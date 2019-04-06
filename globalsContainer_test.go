package staticModel

import "testing"

func TestConfigContainer(t *testing.T) {

	c := &globalsContainer{
		"domain",
		"basePath",
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
		"disqusId",
		"targetDir",
		"description",
		"homeText",
		"homeHeadline",
		"svgLogo",
		"keyWords",
		"subject",
		"author"}

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
