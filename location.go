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
