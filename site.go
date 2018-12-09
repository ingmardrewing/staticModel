package staticModel

import (
	"path"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
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

func RewriteJson(config staticPersistence.JsonConfig) {
	siteCreator := new(siteCreator)
	siteCreator.init(config)
	siteCreator.addConfigData()
	siteCreator.rewritePages()
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
			fl.FileName,
			"")
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
			fl.FileName,
			"")
		s.site.AddMarginal(l)
	}
}

func (s *siteCreator) rewritePages() {
	srcs := s.config.Src
	for _, src := range srcs {
		//dtos := staticPersistence.ReadPagesFromDir(src.Dir)
		staticPersistence.ReadPagesFromDir(src.Dir)
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
			p := NewPage(dto, s.config.Domain, s.site)

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
			emptyPage := NewPage(dto, s.config.Domain, s.site)
			container.AddPage(emptyPage)
		}
		if src.Type == "blog" {
			bnpg := NewBlogNaviPageGenerator(s.site, "/"+src.SubDir, container)
			n := bnpg.Createpages()
			for _, p := range n {
				container.AddNaviPage(p)
			}
		}

		if src.Type == "blog" || src.Type == "narrative" {
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
		if src.Type == "portfolio" {
			pages := container.Pages()
			for _, pg := range pages {
				container.AddRepresentational(pg)
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
