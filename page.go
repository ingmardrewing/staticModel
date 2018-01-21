package staticModel

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
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
func NewMarginalPage(dto staticPersistence.DTO) staticIntf.Page {
	page := new(marginalPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return fillPage(page, dto)
}

// NewPostPage
func NewPostPage(dto staticPersistence.DTO) staticIntf.Page {
	page := new(postPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return fillPage(page, dto)
}

// NewNaviPage
func NewNaviPage(dto staticPersistence.DTO) staticIntf.Page {
	page := new(naviPage)
	page.doc = htmlDoc.NewHtmlDoc()
	return fillPage(page, dto)
}

// NewPage
func fillPage(page staticIntf.Page, dto staticPersistence.DTO) staticIntf.Page {
	page.Domain(dto.Domain())
	page.Title(dto.Title())
	page.ThumbnailUrl(dto.ThumbUrl())
	page.Id(dto.Id())
	page.Description(dto.Description())
	page.Content(dto.Content())
	page.ImageUrl(dto.ImageUrl())
	page.PublishedTime(dto.CreateDate())
	page.DisqusId(dto.DisqusId())

	page.Filename(dto.Filename())
	page.PathFromDocRoot(dto.PathFromDocRoot())
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
	return " "
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
