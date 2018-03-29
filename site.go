package staticModel

import (
	"path"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
	"github.com/ingmardrewing/staticUtil"
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
}

func (s *siteCreator) addLocations() {

	// add configured main navigation
	for _, fl := range s.config.Context.Header {
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
	for _, fl := range s.config.Context.Footer {
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
	if staticUtil.DirExists(s.config.Src.PostsDir) {
		dtos := staticPersistence.ReadPagesFromDir(s.config.Src.PostsDir)
		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)
			newPath := path.Join("/blog/", p.PathFromDocRoot())
			p.PathFromDocRoot(newPath)
			s.site.addPost(p)
		}
		bnpg := NewBlogNaviPageGenerator(s.site, "/blog/")
		s.site.postNaviPages = bnpg.Createpages()
	}

	if staticUtil.DirExists(s.config.Src.MainPages) {
		dtos := staticPersistence.ReadPagesFromDir(s.config.Src.MainPages)
		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)
			s.site.addMainPage(p)
		}
	}

	if staticUtil.DirExists(s.config.Src.MarginalDir) {
		dtos := staticPersistence.ReadPagesFromDir(s.config.Src.MarginalDir)
		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)
			s.site.addMarginalPage(p)
		}

		locs := ElementsToLocations(s.site.Marginals())
		for _, l := range locs {
			s.site.AddMarginal(l)
		}
	}

	if staticUtil.DirExists(s.config.Src.Narrative) {
		dtos := staticPersistence.ReadPagesFromDir(s.config.Src.Narrative)
		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)
			s.site.addNarrativePage(p)
		}
	}

	if staticUtil.DirExists(s.config.Src.NarrativeMarginals) {
		dtos := staticPersistence.ReadPagesFromDir(
			s.config.Src.NarrativeMarginals)
		for _, dto := range dtos {
			p := NewPage(dto, s.config.Domain)
			s.site.addNarrativeMarginalPage(p)
		}
	}
}

type siteDto struct {
	pagesContainer
	locationsContainer
	configContainer
}

type pagesContainer struct {
	posts                  []staticIntf.Page
	postNaviPages          []staticIntf.Page
	mainPages              []staticIntf.Page
	marginalPages          []staticIntf.Page
	narrativePages         []staticIntf.Page
	narrativeMarginalPages []staticIntf.Page
	narrativeArchivePages  []staticIntf.Page
}

func (c *pagesContainer) Posts() []staticIntf.Page {
	return c.posts
}

func (c *pagesContainer) PostNaviPages() []staticIntf.Page {
	return c.postNaviPages
}

func (c *pagesContainer) Pages() []staticIntf.Page {
	return c.mainPages
}

func (c *pagesContainer) Marginals() []staticIntf.Page {
	return c.marginalPages
}

func (c *pagesContainer) Narratives() []staticIntf.Page {
	return c.narrativePages
}

func (c *pagesContainer) NarrativeMarginals() []staticIntf.Page {
	return c.narrativeMarginalPages
}

func (c *pagesContainer) addMainPage(p staticIntf.Page) {
	c.mainPages = append(c.mainPages, p)
}

func (c *pagesContainer) addMarginalPage(p staticIntf.Page) {
	c.marginalPages = append(c.marginalPages, p)
}

func (c *pagesContainer) addPost(p staticIntf.Page) {
	c.posts = append(c.posts, p)
}

func (c *pagesContainer) addNarrativePage(p staticIntf.Page) {
	c.narrativePages = append(c.narrativePages, p)
}

func (c *pagesContainer) addNarrativeMarginalPage(p staticIntf.Page) {
	c.narrativeMarginalPages = append(c.narrativeMarginalPages, p)
}

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
}

func (s *configContainer) TwitterHandle() string { return s.twitterHandle }

func (s *configContainer) Description() string { return s.description }

func (s *configContainer) Topic() string { return s.topic }

func (s *configContainer) Tags() string { return s.tags }

func (s *configContainer) Site() string { return s.domain }

func (s *configContainer) CardType() string { return s.cardType }

func (s *configContainer) Section() string { return s.section }

func (s *configContainer) FBPage() string { return s.fbPage }

func (s *configContainer) TwitterPage() string { return s.twitterPage }

func (s *configContainer) RssPath() string { return s.rssPath }

func (s *configContainer) RssFilename() string { return s.rssFilename }

func (s *configContainer) Css() string { return s.css }

func (s *configContainer) Domain() string { return s.domain }

func (s *configContainer) DisqusId() string { return s.disqusId }

func (s *configContainer) TargetDir() string { return s.targetDir }
