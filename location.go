package staticModel

// Creates a new Loc, implementing Location Interface
func NewLocation(url, prodDomain, title, thumbnailUrl, fsPath, fsFilename string) *loc {
	return &loc{url, prodDomain, title, thumbnailUrl, fsPath, fsFilename}
}

type loc struct {
	url             string
	domain          string
	title           string
	thumbnailUrl    string
	pathFromDocRoot string
	htmlfilename    string
}

func (l *loc) Domain(domain ...string) string {
	if len(domain) > 0 {
		l.domain = domain[0]
	}
	return l.domain
}

func (l *loc) PathFromDocRoot(pathFromDocRoot ...string) string {
	if len(pathFromDocRoot) > 0 {
		l.pathFromDocRoot = pathFromDocRoot[0]
	}
	return l.pathFromDocRoot
}

func (l *loc) HtmlFilename(htmlfilename ...string) string {
	if len(htmlfilename) > 0 {
		l.htmlfilename = htmlfilename[0]
	}
	return l.htmlfilename
}

func (l *loc) Url(url ...string) string {
	if len(url) > 0 {
		l.url = url[0]
	}
	return l.url
}

func (l *loc) Title(title ...string) string {
	if len(title) > 0 {
		l.title = title[0]
	}
	return l.title
}

func (l *loc) ThumbnailUrl(thumbnailUrl ...string) string {
	if len(thumbnailUrl) > 0 {
		l.thumbnailUrl = thumbnailUrl[0]
	}
	return l.thumbnailUrl
}
