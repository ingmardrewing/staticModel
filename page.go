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

	// TODO: Make page understand image splice
	if len(dto.Images()) > 0 {
		firstImage := dto.Images()[0]
		page.microThumbnailUrl = firstImage.W185Square()
		page.thumbnailUrl = firstImage.W390Square()
		page.imageUrl = firstImage.W800Square()

		// TODO: Reimplement narrative archive page
		// and remove this completely:
		page.thumbBase64 = ""
	}

	page.images = dto.Images()
	page.title = dto.Title()
	page.description = dto.Description()
	page.content = dto.Content()
	page.category = dto.Category()
	page.publishedTime = dto.CreateDate()
	page.htmlfilename = dto.Filename()
	page.pathFromDocRoot = dto.PathFromDocRoot()
	page.pathFromDocRootWithName = path.Join(dto.PathFromDocRoot(), dto.Filename())
	page.url = "https://" + path.Join(domain, dto.PathFromDocRoot(), dto.Filename())

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
