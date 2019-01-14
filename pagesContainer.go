package staticModel

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewPagesContainer(variant, headline string) staticIntf.PagesContainer {
	pc := new(pagesContainer)
	pc.variant = variant
	pc.headline = headline
	pc.pages = []staticIntf.Page{}
	pc.naviPages = []staticIntf.Page{}
	pc.representationals = []staticIntf.Page{}
	return pc
}

type pagesContainer struct {
	variant           string
	headline          string
	pages             []staticIntf.Page
	naviPages         []staticIntf.Page
	representationals []staticIntf.Page
}

func (c *pagesContainer) Variant() string {
	return c.variant
}

func (c *pagesContainer) Headline() string {
	return c.headline
}

func (c *pagesContainer) Representationals() []staticIntf.Page {
	return c.representationals
}

func (c *pagesContainer) Pages() []staticIntf.Page {
	return c.pages
}

func (c *pagesContainer) NaviPages() []staticIntf.Page {
	return c.naviPages
}

func (c *pagesContainer) AddRepresentational(p staticIntf.Page) {
	c.representationals = append(c.representationals, p)
}

func (c *pagesContainer) AddPage(p staticIntf.Page) {
	tool := staticUtil.NewPagesContainerTool(c)
	if tool.GetIndexOfPage(p) == -1 {
		c.pages = append(c.pages, p)
	}
}

func (c *pagesContainer) AddNaviPage(p staticIntf.Page) {
	tool := staticUtil.NewPagesContainerTool(c)
	if tool.GetIndexOfNaviPage(p) == -1 {
		c.naviPages = append(c.naviPages, p)
	}
}
