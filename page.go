package staticModel

import (
	"path"
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

func NewPage(
	dto staticIntf.PageDto,
	domain string,
	site staticIntf.Site) *page {

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
	page.htmlfilename = dto.HtmlFilename()
	page.thumbBase64 = dto.ThumbBase64()
	page.pathFromDocRoot = dto.PathFromDocRoot()
	page.pathFromDocRootWithName = path.Join(dto.PathFromDocRoot(), dto.HtmlFilename())
	page.url = "https://" + path.Join(domain, dto.PathFromDocRoot(), dto.HtmlFilename())

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
	if p.site != nil {
		l := "/" + path.Join(p.site.BasePath(), p.pathFromDocRoot, p.htmlfilename)
		if strings.HasPrefix(l, "//") {
			l = strings.TrimPrefix(l, "/")
		}
		return l
	}
	l := "/" + path.Join(p.pathFromDocRoot, p.htmlfilename)
	if strings.HasPrefix(l, "//") {
		l = strings.TrimPrefix(l, "/")
	}
	return l
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
