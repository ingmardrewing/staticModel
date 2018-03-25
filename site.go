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

	site := new(siteDto)

	site.twitterHandle = config.Context.TwitterHandle
	site.topic = config.Context.Topic
	site.tags = config.Context.Tags
	site.domain = config.Domain
	site.cardType = config.Context.CardType
	site.section = config.Context.Section
	site.fbPage = config.Context.FbPage
	site.twitterPage = config.Context.TwitterPage
	site.rssPath = config.Deploy.RssPath
	site.rssFilename = config.Deploy.RssFilename
	site.css = config.Deploy.CssFileName
	site.disqusId = config.Context.DisqusShortname
	site.targetDir = config.Deploy.TargetDir

	site.description = config.DefaultMeta.BlogExcerpt

	if staticUtil.DirExists(config.Src.PostsDir) {
		dtos := staticPersistence.ReadPosts(config.Src.PostsDir)
		for _, dto := range dtos {
			p := NewPage(dto, config.Domain)
			newPath := path.Join("/blog/", p.PathFromDocRoot())
			p.PathFromDocRoot(newPath)
			site.addPost(p)
		}
		bnpg := NewBlogNaviPageGenerator(site, "/blog/")
		site.postNaviPages = bnpg.Createpages()
	}

	if staticUtil.DirExists(config.Src.MainPages) {
		dtos := staticPersistence.ReadPages(config.Src.MainPages)
		for _, dto := range dtos {
			p := NewPage(dto, config.Domain)
			site.addMainPage(p)
		}
	}

	if staticUtil.DirExists(config.Src.MarginalDir) {
		dtos := staticPersistence.ReadMarginals(config.Src.MarginalDir)
		for _, dto := range dtos {
			p := NewPage(dto, config.Domain)
			site.addMarginalPage(p)
		}

		locs := ElementsToLocations(site.Marginals())
		for _, l := range locs {
			site.AddMarginal(l)
		}
	}

	if staticUtil.DirExists(config.Src.Narrative) {
		dtos := staticPersistence.ReadNarrativePages(config.Src.Narrative)
		for _, dto := range dtos {
			p := NewPage(dto, config.Domain)
			site.addNarrativePage(p)
		}
	}

	if staticUtil.DirExists(config.Src.NarrativeMarginals) {
		dtos := staticPersistence.ReadMarginals(
			config.Src.NarrativeMarginals)
		for _, dto := range dtos {
			p := NewPage(dto, config.Domain)
			site.addNarrativeMarginalPage(p)
		}
	}

	// add configured main navigation
	for _, fl := range config.Context.Header {
		l := NewLocation(
			fl.ExternalLink,
			config.Domain,
			fl.Label,
			"",
			fl.Path,
			fl.FileName)
		site.AddMain(l)
	}

	// add configured marginal navigation
	for _, fl := range config.Context.Footer {
		l := NewLocation(
			fl.ExternalLink,
			config.Domain,
			fl.Label,
			"",
			fl.Path,
			fl.FileName)
		site.AddMarginal(l)
	}

	return site
}

type siteDto struct {
	main, marginal         []staticIntf.Location
	posts                  []staticIntf.Page
	postNaviPages          []staticIntf.Page
	mainPages              []staticIntf.Page
	marginalPages          []staticIntf.Page
	narrativePages         []staticIntf.Page
	narrativeMarginalPages []staticIntf.Page
	narrativeArchivePages  []staticIntf.Page

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

func (c *siteDto) Posts() []staticIntf.Page {
	return c.posts
}

func (c *siteDto) PostNaviPages() []staticIntf.Page {
	return c.postNaviPages
}

func (c *siteDto) Pages() []staticIntf.Page {
	return c.mainPages
}

func (c *siteDto) Marginals() []staticIntf.Page {
	return c.marginalPages
}

func (c *siteDto) Narratives() []staticIntf.Page {
	return c.narrativePages
}

func (c *siteDto) NarrativeMarginals() []staticIntf.Page {
	return c.narrativeMarginalPages
}

func (c *siteDto) addMainPage(p staticIntf.Page) {
	c.mainPages = append(c.mainPages, p)
}

func (c *siteDto) addMarginalPage(p staticIntf.Page) {
	c.marginalPages = append(c.marginalPages, p)
}

func (c *siteDto) addPost(p staticIntf.Page) {
	c.posts = append(c.posts, p)
}

func (c *siteDto) addNarrativePage(p staticIntf.Page) {
	c.narrativePages = append(c.narrativePages, p)
}

func (c *siteDto) addNarrativeMarginalPage(p staticIntf.Page) {
	c.narrativeMarginalPages = append(c.narrativeMarginalPages, p)
}

func (c *siteDto) AddMain(loc staticIntf.Location) {
	c.add(&c.main, loc)
}

func (c *siteDto) Main() []staticIntf.Location {
	return c.main
}

func (c *siteDto) AddMarginal(loc staticIntf.Location) {
	c.add(&c.marginal, loc)
}

func (c *siteDto) Marginal() []staticIntf.Location {
	return c.marginal
}

func (c *siteDto) add(collection *[]staticIntf.Location, locs ...staticIntf.Location) {
	for _, l := range locs {
		*collection = append(*collection, l)
	}
}

func (s *siteDto) TwitterHandle() string { return s.twitterHandle }

func (s *siteDto) Description() string { return s.description }

func (s *siteDto) Topic() string { return s.topic }

func (s *siteDto) Tags() string { return s.tags }

func (s *siteDto) Site() string { return s.domain }

func (s *siteDto) CardType() string { return s.cardType }

func (s *siteDto) Section() string { return s.section }

func (s *siteDto) FBPage() string { return s.fbPage }

func (s *siteDto) TwitterPage() string { return s.twitterPage }

func (s *siteDto) RssPath() string { return s.rssPath }

func (s *siteDto) RssFilename() string { return s.rssFilename }

func (s *siteDto) Css() string { return s.css }

func (s *siteDto) Domain() string { return s.domain }

func (s *siteDto) DisqusId() string { return s.disqusId }

func (s *siteDto) TargetDir() string { return s.targetDir }
