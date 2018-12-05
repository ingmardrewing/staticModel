package staticModel

import (
	"regexp"
	"strconv"
	"time"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// NewMarginalPage
func NewPage(dto staticIntf.PageDto, domain string) staticIntf.Page {
	page := new(page)
	page.doc = htmlDoc.NewHtmlDoc()
	page.domain = domain
	fillPage(page, dto)
	return page
}

func fillPage(page *page, dto staticIntf.PageDto) staticIntf.Page {
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
	page.pathFromDocRoot = dto.PathFromDocRoot()
	page.thumbBase64 = dto.ThumbBase64()
	return page
}

// page
type page struct {
	loc
	pageContent
	site           staticIntf.Site
	navigatedPages []staticIntf.Page
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

func (p *page) Site(s ...staticIntf.Site) staticIntf.Site {
	if len(s) == 1 {
		p.site = s[0]
	}
	return p.site
}

// pageContent
type pageContent struct {
	doc           *htmlDoc.HtmlDoc
	id            int
	content       string
	description   string
	imageUrl      string
	publishedTime string
	disqusId      string
	thumbBase64   string
	category      string
}

func (p pageContent) GetDoc() *htmlDoc.HtmlDoc { return p.doc }

func (p pageContent) Category() string { return p.category }

func (p pageContent) Id() int { return p.id }

func (p pageContent) ThumbBase64() string { return p.thumbBase64 }

func (p pageContent) DisqusId() string { return p.disqusId }

func (p pageContent) Content() string { return p.content }

func (p pageContent) Description() string { return p.description }

func (p pageContent) ImageUrl() string { return p.imageUrl }

func (p pageContent) PublishedTime(format ...string) string {
	rx := regexp.MustCompile("(\\d{4})-(\\d{1,2})-(\\d{1,2}) (\\d{1,2}):(\\d{1,2}):(\\d{1,2})")
	m := rx.FindStringSubmatch(p.publishedTime)

	if len(m) > 1 {
		m := rx.FindStringSubmatch(p.publishedTime)
		conv := func(a string) int { i, _ := strconv.Atoi(a); return i }
		loc, _ := time.LoadLocation("Europe/Berlin")
		t := time.Date(
			conv(m[1]),
			time.Month(conv(m[2])),
			conv(m[3]),
			conv(m[4]),
			conv(m[5]),
			conv(m[6]),
			0,
			loc)
		if len(format) > 0 {
			return t.Format(format[0])
		}
		stamp := t.Format("Mon, 02 Jan 2006 15:04:05")
		return stamp + " +0100"
	}

	return p.publishedTime
}
