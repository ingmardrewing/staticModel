package staticModel

import (
	"testing"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
)

func TestNewPage(t *testing.T) {
	p := getPage("")

	expected := p.Id()
	actual := 42

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	actualString := p.Category()
	expectedString := "categoryValue"

	if actualString != expectedString {
		t.Error("Expected", expectedString, "but got", actualString)
	}

	actualDate := "createDateValue"
	expectedDate := p.PublishedTime()

	if actualDate != expectedDate {
		t.Error("Expected", expectedDate, "but got", actualDate)
	}
}

func TestAddHeaderNodes(t *testing.T) {
	p := getPage("")

	n := htmlDoc.NewNode("link", "", "src", "/test")

	p.AddHeaderNodes([]*htmlDoc.Node{n})

	actual := len(p.doc.Render())
	expected := 73

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestAddBodyNodes(t *testing.T) {
	p := getPage("")

	n := htmlDoc.NewNode("a", "", "src", "/test/asdf")

	p.AddBodyNodes([]*htmlDoc.Node{n})
	doc := p.GetDoc()

	actual := len(doc.Render())
	expected := 78

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestImageUrl(t *testing.T) {
	p := getPage("")

	actual := p.ImageUrl()
	expected := "imageUrlValue"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDescription(t *testing.T) {
	p := getPage("")

	actual := p.Description()
	expected := "descriptionValue"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestContent(t *testing.T) {
	p := getPage("")

	actual := p.Content()
	expected := "contentValue"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDisqusId(t *testing.T) {
	p := getPage("")

	actual := p.DisqusId()
	expected := "disqusIdValue"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestThumbBase64(t *testing.T) {
	p := getPage("")

	actual := p.ThumbBase64()
	expected := "thumbBase64Value"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestAcceptVisitor(t *testing.T) {
	p := getPage("")
	v := new(mockedComponent)
	p.AcceptVisitor(v)

	if !v.visitPageCalled {
		t.Error("Expected Visitpage to be called, but it wasn't.")
	}
}

func TestReformatedPublishedTime(t *testing.T) {
	p := getPage("2018-3-14 23:13:58")

	actual := p.PublishedTime()
	expected := "Wed, 14 Mar 2018 23:13:58 +0100"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func getPage(time string) *page {
	t := "createDateValue"
	if len(time) > 0 {
		t = time
	}
	dto := staticPersistence.NewFilledDto(42,
		"titleValue",
		"titlePlainValue",
		"thumbUrlValue",
		"imageUrlValue",
		"descriptionValue",
		"disqusIdValue",
		t,
		"contentValue",
		"urlValue",
		"domainValue",
		"pathValue",
		"fspathValue",
		"htmlfilenameValue",
		"thumbBase64Value",
		"categoryValue")

	page := new(page)
	page.doc = htmlDoc.NewHtmlDoc()
	page.domain = "testDomain"
	fillPage(page, dto)
	return page
}

type mockedComponent struct {
	visitPageCalled bool
}

func (m *mockedComponent) VisitPage(p staticIntf.Page) { m.visitPageCalled = true }

func (m *mockedComponent) GetCss() string { return "" }

func (m *mockedComponent) GetJs() string { return "" }

func (m *mockedComponent) Renderer(r ...staticIntf.Renderer) staticIntf.Renderer { return nil }
