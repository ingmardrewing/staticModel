package staticModel

import (
	"path"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
	log "github.com/sirupsen/logrus"
)

// Creates a site dto and the  pages
// and links being parts of the site in turn
func NewSiteDto(config staticPersistence.JsonConfig) staticIntf.Site {
	siteCreator := new(siteCreator)
	siteCreator.init(config)
	siteCreator.addConfigData()
	siteCreator.addPages()
	siteCreator.addLocations()
	return siteCreator.site
}

type siteCreator struct {
	site   *siteDto
	config staticPersistence.JsonConfig
}

func (s *siteCreator) init(config staticPersistence.JsonConfig) {
	s.config = config
	s.site = new(siteDto)
}

func (s *siteCreator) addConfigData() {
	s.site.twitterHandle = s.config.Context.TwitterHandle
	s.site.topic = s.config.Context.Topic
	s.site.tags = s.config.Context.Tags
	s.site.domain = s.config.Domain
	s.site.cardType = s.config.Context.CardType
	s.site.section = s.config.Context.Section
	s.site.fbPage = s.config.Context.FbPage
	s.site.twitterPage = s.config.Context.TwitterPage
	s.site.rssPath = s.config.Deploy.RssPath
	s.site.rssFilename = s.config.Deploy.RssFilename
	s.site.css = s.config.Deploy.CssFileName
	s.site.disqusId = s.config.Context.DisqusShortname
	s.site.targetDir = s.config.Deploy.TargetDir
	s.site.description = s.config.DefaultMeta.BlogExcerpt
	s.site.homeText = s.config.HomeText
	s.site.homeHeadline = s.config.HomeHeadline
}

func (s *siteCreator) addLocations() {

	// add configured main navigation
	for _, fl := range s.config.Context.MainLinks {
		l := NewLocation(
			fl.ExternalLink,
			s.config.Domain,
			fl.Label,
			"",
			fl.Path,
			fl.FileName)
		s.site.AddMain(l)
	}

	// add configured marginal navigation
	for _, fl := range s.config.Context.MarginalLinks {
		l := NewLocation(
			fl.ExternalLink,
			s.config.Domain,
			fl.Label,
			"",
			fl.Path,
			fl.FileName)
		s.site.AddMarginal(l)
	}
}

func (s *siteCreator) addPages() {
	srcs := s.config.Src
	for _, src := range srcs {
		dtos := staticPersistence.ReadPagesFromDir(src.Dir)

		container := new(pagesContainer)
		container.variant = src.Type
		s.site.AddContainer(container)

		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)

			newPath := path.Join(src.SubDir, p.PathFromDocRoot())
			p.PathFromDocRoot(newPath)

			container.AddPage(p)
		}

		if src.Type == "main" {
			dto := staticPersistence.NewFilledDto(
				0,
				s.config.Domain,
				s.config.Domain,
				"",
				"",
				s.config.DefaultMeta.BlogExcerpt,
				"",
				"",
				"",
				"",
				s.config.Domain,
				"/",
				"/",
				"index.html",
				"",
				"main",
				"")
			emptyPage := NewPage(dto, s.config.Domain)
			container.AddPage(emptyPage)
		}
		if src.Type == "blog" {
			bnpg := NewBlogNaviPageGenerator(s.site, "/"+src.SubDir, container)
			n := bnpg.Createpages()
			for _, p := range n {
				container.AddNaviPage(p)
			}
		}

		if src.Type == "blog" || src.Type == "narrative" || src.Type == "portfolio" {
			pages := container.Pages()
			nrOfRepPages := 4
			if src.Type == "portfolio" {
				nrOfRepPages = len(pages) - 1
			}
			if len(pages) > nrOfRepPages {
				for _, pg := range pages[len(pages)-nrOfRepPages:] {
					container.AddRepresentational(pg)
				}
			}
		}
		if src.Type == "marginal" {
			locs := ElementsToLocations(container.Pages())
			for _, l := range locs {
				s.site.AddMarginal(l)
			}
		}
	}
}

type siteDto struct {
	pagesContainerCollection
	locationsContainer
	configContainer
}

type pagesContainer struct {
	variant           string
	pages             []staticIntf.Page
	naviPages         []staticIntf.Page
	representationals []staticIntf.Page
}

func (c *pagesContainer) Variant() string {
	return c.variant
}

func (c *pagesContainer) Representationals() []staticIntf.Page {
	return c.representationals
}

func (c *pagesContainer) Pages() []staticIntf.Page {
	return c.pages
}

func (c *pagesContainer) NaviPages() []staticIntf.Page {
	return c.naviPages
}

func (c *pagesContainer) AddRepresentational(p staticIntf.Page) {
	c.representationals = append(c.representationals, p)
}

func (c *pagesContainer) AddPage(p staticIntf.Page) {
	c.pages = append(c.pages, p)
}

func (c *pagesContainer) AddNaviPage(p staticIntf.Page) {
	c.naviPages = append(c.naviPages, p)
}

//
type pagesContainerCollection struct {
	containers []staticIntf.PagesContainer
}

func (c *pagesContainerCollection) AddContainer(p staticIntf.PagesContainer) {
	c.containers = append(c.containers, p)
}

func (c *pagesContainerCollection) Containers() []staticIntf.PagesContainer {
	return c.containers
}

func (c *pagesContainerCollection) ContainersOrderedByVariants(variants ...string) []staticIntf.PagesContainer {
	log.Debug("ContainersOrderedByVariants, nr of containers:", len(c.containers))
	orderedContainers := []staticIntf.PagesContainer{}
	for _, v := range variants {
		log.Debug("ContainersOrderedByVariants - looping through variant:", v)
		container := c.getContainerByVariant(v)
		if container != nil {
			orderedContainers = append(orderedContainers, container)
		}
	}
	return orderedContainers
}

func (c *pagesContainerCollection) getContainerByVariant(v string) staticIntf.PagesContainer {
	for _, co := range c.containers {
		log.Debug("getContainerByVariant: ", co.Variant(), "==?", v)
		if co.Variant() == v {
			log.Debug("getContainerByVariant, returning: ", co.Variant())
			return co
		}
	}
	return nil
}

func (c *pagesContainerCollection) getPagesByVariant(v string, navi bool) []staticIntf.Page {
	co := c.getContainerByVariant(v)
	if co != nil {
		if navi {
			return co.NaviPages()
		} else {
			return co.Pages()
		}
	}
	return nil
}

func (c *pagesContainerCollection) Pages() []staticIntf.Page {
	return c.getPagesByVariant("pages", false)
}

func (c *pagesContainerCollection) Home() []staticIntf.Page {
	pp := c.getPagesByVariant("main", false)
	return pp
}

func (c *pagesContainerCollection) Portfolio() []staticIntf.Page {
	pp := c.getPagesByVariant("portfolio", false)
	return pp
}

func (c *pagesContainerCollection) Posts() []staticIntf.Page {
	pp := c.getPagesByVariant("blog", false)
	return pp
}

func (c *pagesContainerCollection) PostNaviPages() []staticIntf.Page {
	return c.getPagesByVariant("blog", true)
}

func (c *pagesContainerCollection) Marginals() []staticIntf.Page {
	return c.getPagesByVariant("marginal", false)
}

func (c *pagesContainerCollection) Narratives() []staticIntf.Page {
	return c.getPagesByVariant("narrative", false)
}

func (c *pagesContainerCollection) NarrativeMarginals() []staticIntf.Page {
	return c.getPagesByVariant("narrativesMarginals", false)
}

//
type locationsContainer struct {
	main, marginal []staticIntf.Location
}

func (c *locationsContainer) AddMain(loc staticIntf.Location) {
	c.add(&c.main, loc)
}

func (c *locationsContainer) Main() []staticIntf.Location {
	return c.main
}

func (c *locationsContainer) AddMarginal(loc staticIntf.Location) {
	c.add(&c.marginal, loc)
}

func (c *locationsContainer) Marginal() []staticIntf.Location {
	return c.marginal
}

func (c *locationsContainer) add(collection *[]staticIntf.Location, locs ...staticIntf.Location) {
	for _, l := range locs {
		*collection = append(*collection, l)
	}
}

type configContainer struct {
	twitterHandle string
	topic         string
	tags          string
	cardType      string
	section       string
	fbPage        string
	twitterPage   string
	rssPath       string
	rssFilename   string
	css           string
	domain        string
	disqusId      string
	targetDir     string
	description   string
	homeText      string
	homeHeadline  string
}

func (cc *configContainer) TwitterHandle() string { return cc.twitterHandle }

func (cc *configContainer) Description() string { return cc.description }

func (cc *configContainer) Topic() string { return cc.topic }

func (cc *configContainer) Tags() string { return cc.tags }

func (cc *configContainer) Site() string { return cc.domain }

func (cc *configContainer) CardType() string { return cc.cardType }

func (cc *configContainer) Section() string { return cc.section }

func (cc *configContainer) FBPage() string { return cc.fbPage }

func (cc *configContainer) TwitterPage() string { return cc.twitterPage }

func (cc *configContainer) RssPath() string { return cc.rssPath }

func (cc *configContainer) RssFilename() string { return cc.rssFilename }

func (cc *configContainer) Css() string { return cc.css }

func (cc *configContainer) Domain() string { return cc.domain }

func (cc *configContainer) DisqusId() string { return cc.disqusId }

func (cc *configContainer) TargetDir() string { return cc.targetDir }

func (cc *configContainer) HomeText() string { return cc.homeText }

func (cc *configContainer) HomeHeadline() string { return cc.homeHeadline }
