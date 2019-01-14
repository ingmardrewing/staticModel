package staticModel

import "github.com/ingmardrewing/staticIntf"

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

func (c *pagesContainer) getIndexOfPage(p staticIntf.Page) int {
	return c.getIndex(c.pages, p)
}

func (c *pagesContainer) getIndexOfNaviPage(p staticIntf.Page) int {
	return c.getIndex(c.naviPages, p)
}

func (c *pagesContainer) getIndex(
	pgs []staticIntf.Page,
	p staticIntf.Page) int {

	for i, l := range pgs {
		if l.Link() == p.Link() {
			return i
		}
	}
	return -1
}

func (c *pagesContainer) GetLastPage(p staticIntf.Page) staticIntf.Page {
	if c.getIndexOfPage(p) != -1 {
		return c.getLastOf(c.pages)
	} else if c.getIndexOfNaviPage(p) != -1 {
		return c.getLastOf(c.naviPages)
	}
	return nil
}

func (c *pagesContainer) getLastOf(pgs []staticIntf.Page) staticIntf.Page {
	if len(pgs) > 0 {
		return pgs[len(pgs)-1]
	}
	return nil
}

func (c *pagesContainer) GetFirstPage(p staticIntf.Page) staticIntf.Page {
	if c.getIndexOfPage(p) != -1 {
		return c.getFirstOf(c.pages)
	} else if c.getIndexOfNaviPage(p) != -1 {
		return c.getFirstOf(c.naviPages)
	}
	return nil
}

func (c *pagesContainer) getFirstOf(pgs []staticIntf.Page) staticIntf.Page {

	if len(pgs) > 0 {
		return pgs[0]
	}
	return nil
}

func (c *pagesContainer) GetPageBefore(p staticIntf.Page) staticIntf.Page {
	if c.getIndexOfPage(p) != -1 {
		return c.pageBefore(c.pages, p)
	} else if c.getIndexOfNaviPage(p) != -1 {
		return c.pageBefore(c.naviPages, p)
	}
	return nil
}

func (c *pagesContainer) pageBefore(pgs []staticIntf.Page, p staticIntf.Page) staticIntf.Page {
	i := c.getIndex(pgs, p)
	if i > 0 {
		return pgs[i-1]
	}
	return nil
}

func (c *pagesContainer) GetPageAfter(p staticIntf.Page) staticIntf.Page {
	if c.getIndexOfPage(p) != -1 {
		return c.pageAfter(c.pages, p)
	} else if c.getIndexOfNaviPage(p) != -1 {
		return c.pageAfter(c.naviPages, p)
	}
	return nil
}

func (c *pagesContainer) pageAfter(pgs []staticIntf.Page, p staticIntf.Page) staticIntf.Page {
	i := c.getIndex(pgs, p)
	if i < len(pgs)-1 {
		return pgs[i+1]
	}
	return nil
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

func (c *pagesContainer) SiblingPages(p staticIntf.Page) []staticIntf.Page {
	if c.getIndexOfPage(p) != -1 {
		return c.pages
	} else if c.getIndexOfNaviPage(p) != -1 {
		return c.naviPages
	}
	return nil
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
	if c.getIndexOfPage(p) == -1 {
		c.pages = append(c.pages, p)
		p.Container(c)
	}
}

func (c *pagesContainer) AddNaviPage(p staticIntf.Page) {
	if c.getIndexOfNaviPage(p) == -1 {
		c.naviPages = append(c.naviPages, p)
		p.Container(c)
	}
}
