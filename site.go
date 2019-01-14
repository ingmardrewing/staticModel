package staticModel

func NewSiteDto(
	twitterHandle,
	topic,
	tags,
	domain,
	basePath,
	cardType,
	section,
	fbPage,
	twitterPage,
	rssPath,
	rssFilename,
	css,
	disqusId,
	targetDir,
	description,
	homeText,
	homeHeadline string) *siteDto {

	site := new(siteDto)

	site.twitterHandle = twitterHandle
	site.topic = topic
	site.tags = tags
	site.domain = domain
	site.basePath = basePath
	site.cardType = cardType
	site.section = section
	site.fbPage = fbPage
	site.twitterPage = twitterPage
	site.rssPath = rssPath
	site.rssFilename = rssFilename
	site.css = css
	site.disqusId = disqusId
	site.targetDir = targetDir
	site.description = description
	site.homeText = homeText
	site.homeHeadline = homeHeadline

	return site
}

type siteDto struct {
	pagesContainerCollection
	locationsContainer
	configContainer
}
