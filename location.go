package staticModel

// Creates a new Loc, implementing Location Interface
func NewLocation(
	externalLink,
	domain,
	title,
	thumbnailUrl,
	pathFromDocRoot,
	htmlfilename,
	microThumbnailUrl,
	pathFromDocRootWithName,
	url string) *loc {

	l := new(loc)
	l.extLink = externalLink
	l.domain = domain
	l.title = title
	l.thumbnailUrl = thumbnailUrl
	l.pathFromDocRoot = pathFromDocRoot
	l.htmlfilename = htmlfilename
	l.microThumbnailUrl = microThumbnailUrl
	l.pathFromDocRootWithName = pathFromDocRootWithName
	l.url = url

	return l
}

type loc struct {
	extLink                 string
	domain                  string
	title                   string
	thumbnailUrl            string
	pathFromDocRoot         string
	htmlfilename            string
	microThumbnailUrl       string
	pathFromDocRootWithName string
	url                     string
}

func (l *loc) PathFromDocRootWithName() string {
	return l.pathFromDocRootWithName
}

func (l *loc) Url() string {
	return l.url
}

func (l *loc) ExternalLink() string {
	return l.extLink
}

func (l *loc) Domain() string {
	return l.domain
}

func (l *loc) PathFromDocRoot() string {
	return l.pathFromDocRoot
}

func (l *loc) HtmlFilename() string {
	return l.htmlfilename
}

func (l *loc) Title() string {
	return l.title
}

func (l *loc) ThumbnailUrl() string {
	return l.thumbnailUrl
}

func (l *loc) MicroThumbnailUrl() string {
	return l.microThumbnailUrl
}
