package staticModel

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

type postPage struct {
	page
}

type naviPage struct {
	page
}

type marginalPage struct {
	page
}

// NewMarginalPage
func NewMarginalPage() staticIntf.Page {
	page := new(marginalPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return page
}

// NewPostPage
func NewPostPage() staticIntf.Page {
	page := new(postPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return page
}

// NewNaviPage
func NewNaviPage() staticIntf.Page {
	page := new(naviPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return page
}

type page struct {
	loc
	doc           *htmlDoc.HtmlDoc
	id            int
	content       string
	description   string
	imageUrl      string
	publishedTime string
	disqusId      string
}

func (p *page) Id(id ...int) int {
	if len(id) > 0 {
		p.id = id[0]
	}
	return p.id
}

func (p *page) DisqusId(disqusId ...string) string {
	if len(disqusId) > 0 {
		p.disqusId = disqusId[0]
	}
	return p.disqusId
}

func (p *page) Content(content ...string) string {
	if len(content) > 0 {
		p.content = content[0]
	}
	return p.content
}

func (p *page) Description(description ...string) string {
	if len(description) > 0 {
		p.description = description[0]
	}
	return p.description
}

func (p *page) ImageUrl(imageUrl ...string) string {
	if len(imageUrl) > 0 {
		p.imageUrl = imageUrl[0]
	}
	return p.imageUrl
}

func (p *page) PublishedTime(publishedTime ...string) string {
	if len(publishedTime) > 0 {
		p.publishedTime = publishedTime[0]
	}
	return p.publishedTime
}

func (p *page) GetDoc() *htmlDoc.HtmlDoc {
	return p.doc
}

func (p *page) AcceptVisitor(v staticIntf.Component) {
	v.VisitPage(p)
}

func (p *page) AddHeaderNodes(nodes []*htmlDoc.Node) {
	for _, n := range nodes {
		p.doc.AddHeadNode(n)
	}
}

func (p *page) AddBodyNodes(nodes []*htmlDoc.Node) {
	for _, n := range nodes {
		p.doc.AddBodyNode(n)
	}
}
