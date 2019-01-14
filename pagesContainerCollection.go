package staticModel

import "github.com/ingmardrewing/staticIntf"

type pagesContainerCollection struct {
	containers []staticIntf.PagesContainer
}

func (c *pagesContainerCollection) AddContainer(p staticIntf.PagesContainer) {
	c.containers = append(c.containers, p)
}

func (c *pagesContainerCollection) Containers() []staticIntf.PagesContainer {
	return c.containers
}
