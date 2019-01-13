package staticModel

import (
	"fmt"
	"path"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"

	log "github.com/sirupsen/logrus"
)

//
type source interface {
	generate()
	Container() staticIntf.PagesContainer
	SetData(variant, headline, dir, subDir string, site staticIntf.Site, config staticPersistence.JsonConfig)
}

func NewSource(
	variant, dir, subDir, headline string,
	site staticIntf.Site,
	config staticPersistence.JsonConfig) source {

	var s source
	switch variant {
	case staticIntf.HOME:
		s = new(homeSource)
	case staticIntf.BLOG:
		s = new(blogSource)
	case staticIntf.PORTFOLIO:
		s = new(portfolioSource)
	case staticIntf.MARGINALS:
		s = new(marginalSource)
	case staticIntf.NARRATIVES:
		s = new(narrativeSource)
	case staticIntf.NARRATIVEMARGINALS:
		s = new(narrativeMarginalSource)
	default:
		log.Fatal("unaccounted variant:", variant)
	}

	s.SetData(variant, headline, dir, subDir, site, config)

	return s
}

type blogSource struct {
	defaultSource
}

func (bs *blogSource) generate() {
	bs.generateContainer()

	bnpg := NewBlogNaviPageGenerator(
		bs.site,
		"/"+bs.subDir,
		bs.container)
	naviPages := bnpg.Createpages()
	for _, p := range naviPages {
		bs.container.AddNaviPage(p)
	}
	pages := bs.container.Pages()
	nrOfRepPages := 4
	if len(pages) > nrOfRepPages {
		for _, pg := range pages[len(pages)-nrOfRepPages:] {
			bs.container.AddRepresentational(pg)
		}
	}
}

//
type portfolioSource struct {
	defaultSource
}

func (ps *portfolioSource) generate() {
	ps.generateContainer()

	pages := ps.container.Pages()
	for _, pg := range pages {
		ps.container.AddRepresentational(pg)
	}
}

//
type homeSource struct {
	defaultSource
}

func (hs *homeSource) generate() {
	hs.generateContainer()
	dto := staticPersistence.NewFilledDto(
		0,
		hs.config.Domain,
		hs.config.Domain,
		"",
		"",
		hs.config.DefaultMeta.BlogExcerpt,
		"",
		"",
		"",
		"",
		hs.config.Domain,
		"/",
		"/",
		"index.html",
		"",
		"main",
		"")
	homePage := NewPage(dto, hs.config.Domain, hs.site)
	hs.container.AddPage(homePage)
}

//
type narrativeMarginalSource struct {
	defaultSource
}

func (nms *narrativeMarginalSource) generate() {
	nms.generateContainer()
}

//
type marginalSource struct {
	defaultSource
}

func (mrs *marginalSource) generate() {
	mrs.generateContainer()
	locs := ElementsToLocations(mrs.container.Pages())
	for _, l := range locs {
		mrs.site.AddMarginal(l)
	}

}

//
type narrativeSource struct {
	defaultSource
}

func (ns *narrativeSource) generate() {
	ns.generateContainer()

	pages := ns.container.Pages()
	nrOfRepPages := 4
	if len(pages) > nrOfRepPages {
		for _, pg := range pages[len(pages)-nrOfRepPages:] {
			ns.container.AddRepresentational(pg)
		}
	}
}

//
type defaultSource struct {
	variant   string
	headline  string
	dir       string
	subDir    string
	site      staticIntf.Site
	config    staticPersistence.JsonConfig
	container staticIntf.PagesContainer
}

func (a *defaultSource) Container() staticIntf.PagesContainer {
	return a.container
}

func (a *defaultSource) generate() {}

func (a *defaultSource) SetData(variant, headline, dir, subDir string, site staticIntf.Site, config staticPersistence.JsonConfig) {
	a.variant = variant
	a.headline = headline
	a.dir = dir
	a.subDir = subDir
	a.site = site
	a.config = config
}

func (a *defaultSource) generateContainer() {
	log.Debug(fmt.Sprintf("-- new container, type %s, headline %s", a.variant, a.headline))
	container := new(pagesContainer)
	container.variant = a.variant
	container.headline = a.headline
	a.container = container
	for _, dto := range staticPersistence.ReadPagesFromDir(a.dir) {
		a.createPage(dto)
	}
}

func (a *defaultSource) createPage(dto staticIntf.PageDto) {
	p := NewPage(dto, a.config.Domain, a.site)

	newPath := path.Join(a.subDir, p.PathFromDocRoot())
	p.PathFromDocRoot(newPath)

	log.Debug(fmt.Sprintf("page in model, path:%s", newPath))
	a.container.AddPage(p)
}
