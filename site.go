package staticModel

import (
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

type siteDto struct {
	pagesContainerCollection
	locationsContainer
	configContainer
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
	s.site.basePath = s.config.BasePath
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
	for _, src := range s.config.Src {
		//dtos := staticPersistence.ReadPagesFromDir(src.Dir)
		staticPersistence.ReadPagesFromDir(src.Dir)
	}
}

func (s *siteCreator) addPages() {
	for _, src := range s.config.Src {
		sr := NewSource(
			src.Type,
			src.Dir,
			src.SubDir,
			src.Headline,
			s.site,
			s.config)
		sr.generate()
		s.site.AddContainer(sr.Container())
	}
}
