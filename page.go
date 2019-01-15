package staticModel

import (
	"path"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

func NewPage(
	dto staticIntf.PageDto,
	domain string,
	site staticIntf.Site,
	subDir string) *page {

	page := new(page)
	page.doc = htmlDoc.NewHtmlDoc()
	page.domain = domain
	page.site = site

	// fill with data from dto
	page.title = dto.Title()
	page.thumbnailUrl = dto.ThumbUrl()
	page.microThumbnailUrl = dto.MicroThumbUrl()
	page.id = dto.Id()
	page.description = dto.Description()
	page.content = dto.Content()
	page.category = dto.Category()
	page.imageUrl = dto.ImageUrl()
	page.publishedTime = dto.CreateDate()
	page.disqusId = dto.DisqusId()
	page.htmlfilename = dto.HtmlFilename()
	page.pathFromDocRoot = path.Join(subDir, dto.PathFromDocRoot())
	page.thumbBase64 = dto.ThumbBase64()

	return page
}

type page struct {
	loc
	pageContent
	site           staticIntf.Site
	container      staticIntf.PagesContainer
	navigatedPages []staticIntf.Page
}

func (p *page) Container(container ...staticIntf.PagesContainer) staticIntf.PagesContainer {
	if len(container) > 0 {
		p.container = container[0]
	}
	return p.container
}

func (p *page) Link() string {
	return "/" + path.Join(p.site.BasePath(), p.pathFromDocRoot, p.htmlfilename)
}

func (p *page) NavigatedPages(navigatedPages ...staticIntf.Page) []staticIntf.Page {
	if len(navigatedPages) > 0 {
		p.navigatedPages = navigatedPages
	}
	return p.navigatedPages
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

func (p *page) AcceptVisitor(v staticIntf.Component) {
	v.VisitPage(p)
}

func (p *page) Site() staticIntf.Site {
	return p.site
}
