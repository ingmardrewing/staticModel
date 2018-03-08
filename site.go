package staticModel

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
)

func NewSiteDto(config staticPersistence.JsonConfig) staticIntf.Site {

	site := new(siteDto)
	ctxDao := staticPersistence.NewContextDao(config)
	site.contextDto = ctxDao.Dto()

	if len(config.Src.PostsDir) > 0 {
		dtos := staticPersistence.ReadPosts(config.Src.PostsDir)
		for _, dto := range dtos {
			p := NewPostPage(dto)
			site.addPost(p)
		}
	}

	if len(config.Src.MainPages) > 0 {
		dtos := staticPersistence.ReadPages(config.Src.MainPages)
		for _, dto := range dtos {
			p := NewPage(dto)
			site.addMainPage(p)
		}
	}

	if len(config.Src.MarginalDir) > 0 {
		dtos := staticPersistence.ReadMarginals(config.Src.MainPages)
		for _, dto := range dtos {
			p := NewMarginalPage(dto)
			site.addMarginalPage(p)
		}
	}

	if len(config.Src.Narrative) > 0 {
		dtos := staticPersistence.ReadNarrativePages(config.Src.Narrative)
		for _, dto := range dtos {
			p := NewPage(dto)
			site.addNarrativePage(p)
		}
	}

	// add configured main navigation
	for _, fl := range config.Context.Header {
		l := NewLocation(fl.Link, "", fl.Label, "", "", "")
		site.AddMain(l)
	}

	// add configured marginal navigation
	for _, fl := range config.Context.Footer {
		l := NewLocation(fl.Link, "", fl.Label, "", "", "")
		site.AddMarginal(l)
	}

	return site
}

type siteDto struct {
	main, marginal []staticIntf.Location
	contextDto     staticIntf.ContextDto
	posts          []staticIntf.Page
	mainPages      []staticIntf.Page
	marginalPages  []staticIntf.Page
	narrativePages []staticIntf.Page
}

func (c *siteDto) Posts() []staticIntf.Page {
	return c.posts
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

func (c *siteDto) ContextDto() staticIntf.ContextDto {
	return c.contextDto
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
