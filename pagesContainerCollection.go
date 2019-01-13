package staticModel

import (
	"github.com/ingmardrewing/staticIntf"
	log "github.com/sirupsen/logrus"
)

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
		container := c.getContainersByVariant(v)
		if container != nil {
			orderedContainers = append(orderedContainers, container...)
		}
	}
	return orderedContainers
}

func (c *pagesContainerCollection) getContainersByVariant(v string) []staticIntf.PagesContainer {
	containers := []staticIntf.PagesContainer{}
	for _, co := range c.containers {
		if co.Variant() == v {
			containers = append(containers, co)
		}
	}
	return containers
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

func (c *pagesContainerCollection) GetNaviPagesByVariant(v string) []staticIntf.Page {
	nps := []staticIntf.Page{}
	for _, c := range c.getContainersByVariant(v) {
		nps = append(nps, c.NaviPages()...)
	}
	return nps
}

func (c *pagesContainerCollection) GetPagesByVariant(v string) []staticIntf.Page {
	ps := []staticIntf.Page{}
	for _, c := range c.getContainersByVariant(v) {
		ps = append(ps, c.Pages()...)
	}
	return ps
}
