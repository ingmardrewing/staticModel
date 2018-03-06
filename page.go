package staticModel

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

type postPage struct {
	page
}

type marginalPage struct {
	page
}

// NewMarginalPage
func NewPage() staticIntf.Page {
	page := new(page)
	page.doc = htmlDoc.NewHtmlDoc()
	return page
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
func NewNaviPage() staticIntf.NaviPage {
	page := new(naviPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return page
}

type naviPage struct {
	loc
	pageContent
	navigatedPages []staticIntf.Page
}

func (p *naviPage) AcceptVisitor(v staticIntf.Component) {
	v.VisitPage(p)
}

func (np *naviPage) NavigatedPages(navigatedPages ...staticIntf.Page) []staticIntf.Page {
	if len(navigatedPages) > 0 {
		np.navigatedPages = navigatedPages
	}
	return np.navigatedPages
}

type page struct {
	loc
	pageContent
}

func (p *page) AcceptVisitor(v staticIntf.Component) {
	v.VisitPage(p)
}

type pageContent struct {
	doc           *htmlDoc.HtmlDoc
	id            int
	content       string
	description   string
	imageUrl      string
	publishedTime string
	disqusId      string
}

func (p *pageContent) Id(id ...int) int {
	if len(id) > 0 {
		p.id = id[0]
	}
	return p.id
}

func (p *pageContent) DisqusId(disqusId ...string) string {
	if len(disqusId) > 0 {
		p.disqusId = disqusId[0]
	}
	return p.disqusId
}

func (p *pageContent) Content(content ...string) string {
	if len(content) > 0 {
		p.content = content[0]
	}
	return p.content
}

func (p *pageContent) Description(description ...string) string {
	if len(description) > 0 {
		p.description = description[0]
	}
	return p.description
}

func (p *pageContent) ImageUrl(imageUrl ...string) string {
	if len(imageUrl) > 0 {
		p.imageUrl = imageUrl[0]
	}
	return p.imageUrl
}

func (p *pageContent) PublishedTime(publishedTime ...string) string {
	if len(publishedTime) > 0 {
		p.publishedTime = publishedTime[0]
	}
	return p.publishedTime
}

func (p *pageContent) GetDoc() *htmlDoc.HtmlDoc {
	return p.doc
}

func (p *pageContent) AddHeaderNodes(nodes []*htmlDoc.Node) {
	for _, n := range nodes {
		p.doc.AddHeadNode(n)
	}
}

func (p *pageContent) AddBodyNodes(nodes []*htmlDoc.Node) {
	for _, n := range nodes {
		p.doc.AddBodyNode(n)
	}
}
