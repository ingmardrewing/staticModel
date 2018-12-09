package staticModel

import (
	"github.com/ingmardrewing/staticIntf"
	log "github.com/sirupsen/logrus"
)

type pagesContainer struct {
	variant           string
	pages             []staticIntf.Page
	naviPages         []staticIntf.Page
	representationals []staticIntf.Page
}

func (c *pagesContainer) Variant() string {
	return c.variant
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
	c.pages = append(c.pages, p)
}

func (c *pagesContainer) AddNaviPage(p staticIntf.Page) {
	c.naviPages = append(c.naviPages, p)
}

//
type pagesContainerCollection struct {
	containers []staticIntf.PagesContainer
}

func (c *pagesContainerCollection) AddContainer(p staticIntf.PagesContainer) {
	c.containers = append(c.containers, p)
}

func (c *pagesContainerCollection) Containers() []staticIntf.PagesContainer {
	return c.containers
}

func (c *pagesContainerCollection) ContainersOrderedByVariants(variants ...string) []staticIntf.PagesContainer {
	log.Debug("ContainersOrderedByVariants, nr of containers:", len(c.containers))
	orderedContainers := []staticIntf.PagesContainer{}
	for _, v := range variants {
		log.Debug("ContainersOrderedByVariants - looping through variant:", v)
		container := c.getContainerByVariant(v)
		if container != nil {
			orderedContainers = append(orderedContainers, container)
		}
	}
	return orderedContainers
}

func (c *pagesContainerCollection) getContainerByVariant(v string) staticIntf.PagesContainer {
	for _, co := range c.containers {
		log.Debug("getContainerByVariant: ", co.Variant(), "==?", v)
		if co.Variant() == v {
			log.Debug("getContainerByVariant, returning: ", co.Variant())
			return co
		}
	}
	return nil
}

func (c *pagesContainerCollection) getPagesByVariant(v string, navi bool) []staticIntf.Page {
	co := c.getContainerByVariant(v)
	if co != nil {
		if navi {
			return co.NaviPages()
		} else {
			return co.Pages()
		}
	}
	return nil
}

func (c *pagesContainerCollection) Pages() []staticIntf.Page {
	return c.getPagesByVariant("pages", false)
}

func (c *pagesContainerCollection) Home() []staticIntf.Page {
	pp := c.getPagesByVariant("main", false)
	return pp
}

func (c *pagesContainerCollection) Portfolio() []staticIntf.Page {
	pp := c.getPagesByVariant("portfolio", false)
	return pp
}

func (c *pagesContainerCollection) Posts() []staticIntf.Page {
	pp := c.getPagesByVariant("blog", false)
	return pp
}

func (c *pagesContainerCollection) PostNaviPages() []staticIntf.Page {
	return c.getPagesByVariant("blog", true)
}

func (c *pagesContainerCollection) Marginals() []staticIntf.Page {
	return c.getPagesByVariant("marginal", false)
}

func (c *pagesContainerCollection) Narratives() []staticIntf.Page {
	return c.getPagesByVariant("narrative", false)
}

func (c *pagesContainerCollection) NarrativeMarginals() []staticIntf.Page {
	return c.getPagesByVariant("narrativesMarginals", false)
}
