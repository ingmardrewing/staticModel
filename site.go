package staticModel

import "github.com/ingmardrewing/staticIntf"

type Site interface {
	AddMarginalPage(page staticIntf.Page)
	AddNaviPage(page staticIntf.Page)
	AddPost(page staticIntf.Page)
	AddPage(page staticIntf.Page)
	NaviPages() []staticIntf.Page
	Posts() []staticIntf.Page
	Pages() []staticIntf.Page
	MarginalPages() []staticIntf.Page
}

func NewSite() Site {
	return new(site)
}

type site struct {
	rss           string
	javascript    string
	css           string
	marginalPages []staticIntf.Page
	posts         []staticIntf.Page
	naviPages     []staticIntf.Page
	pages         []staticIntf.Page
}

func (s *site) AddPage(page staticIntf.Page) {
	s.pages = append(s.pages, page)
}

func (s *site) AddNaviPage(page staticIntf.Page) {
	s.naviPages = append(s.naviPages, page)
}

func (s *site) AddMarginalPage(page staticIntf.Page) {
	s.marginalPages = append(s.marginalPages, page)
}

func (s *site) AddPost(post staticIntf.Page) {
	s.posts = append(s.posts, post)
}

func (s *site) Pages() []staticIntf.Page {
	return s.pages
}

func (s *site) MarginalPages() []staticIntf.Page {
	return s.marginalPages
}

func (s *site) Posts() []staticIntf.Page {
	return s.posts
}

func (s *site) NaviPages() []staticIntf.Page {
	return s.naviPages
}
