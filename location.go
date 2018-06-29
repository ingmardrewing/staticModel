package staticModel

import (
	"path"
	"strings"
)

// Creates a new Loc, implementing Location Interface
func NewLocation(externalLink, prodDomain, title, thumbnailUrl, pathFromDocRoot, fsFilename, microThumbnailUrl string) *loc {
	return &loc{externalLink, prodDomain, title, thumbnailUrl, pathFromDocRoot, fsFilename, microThumbnailUrl}
}

type loc struct {
	extLink           string
	domain            string
	title             string
	thumbnailUrl      string
	pathFromDocRoot   string
	htmlfilename      string
	microThumbnailUrl string
}

func (l *loc) PathFromDocRootWithName() string {
	p := path.Join(l.pathFromDocRoot, l.htmlfilename)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return p
}

func (l *loc) Url() string {
	p := path.Join(l.domain, l.PathFromDocRootWithName())
	return "https://" + p
}

func (l *loc) ExternalLink(extLink ...string) string {
	if len(extLink) > 0 {
		l.extLink = extLink[0]
	}
	return l.extLink
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

func (l *loc) MicroThumbnailUrl(microThumbnailUrl ...string) string {
	if len(microThumbnailUrl) > 0 {
		l.microThumbnailUrl = microThumbnailUrl[0]
	}
	return l.microThumbnailUrl
}
