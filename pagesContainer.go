package staticModel

import "github.com/ingmardrewing/staticIntf"

type pagesContainer struct {
	variant           string
	headline          string
	pages             []staticIntf.Page
	naviPages         []staticIntf.Page
	representationals []staticIntf.Page
}

func (c *pagesContainer) getIndexOfPage(p staticIntf.Page) int {
	for i, l := range c.pages {
		if l.Link() == p.Link() {
			return i
		}
	}
	return -1
}

func (c *pagesContainer) GetLastPage() staticIntf.Page {
	if len(c.pages) > 0 {
		return c.pages[len(c.pages)-1]
	}
	return nil
}

func (c *pagesContainer) GetFirstPage() staticIntf.Page {
	if len(c.pages) > 0 {
		return c.pages[0]
	}
	return nil
}

func (c *pagesContainer) GetPageBefore(p staticIntf.Page) staticIntf.Page {
	i := c.getIndexOfPage(p)
	if i > 0 {
		return c.pages[i-1]
	}
	return nil
}

func (c *pagesContainer) GetPageAfter(p staticIntf.Page) staticIntf.Page {
	i := c.getIndexOfPage(p)
	if i+1 < len(c.pages) {
		return c.pages[i+1]
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
	if !c.pageAlreadyExists(p) {
		c.pages = append(c.pages, p)
		p.Container(c)
	}
}

func (c *pagesContainer) pageAlreadyExists(p staticIntf.Page) bool {
	for _, otherPage := range c.pages {
		if p.Link() == otherPage.Link() {
			return true
		}
	}
	return false
}

func (c *pagesContainer) AddNaviPage(p staticIntf.Page) {
	c.naviPages = append(c.naviPages, p)
}
