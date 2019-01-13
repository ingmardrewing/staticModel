package staticModel

import "github.com/ingmardrewing/staticIntf"

//
type locationsContainer struct {
	main, marginal []staticIntf.Location
}

func (c *locationsContainer) AddMain(loc staticIntf.Location) {
	c.add(&c.main, loc)
}

func (c *locationsContainer) Main() []staticIntf.Location {
	return c.main
}

func (c *locationsContainer) AddMarginal(loc staticIntf.Location) {
	c.add(&c.marginal, loc)
}

func (c *locationsContainer) Marginal() []staticIntf.Location {
	return c.marginal
}

func (c *locationsContainer) add(collection *[]staticIntf.Location, locs ...staticIntf.Location) {
	for _, l := range locs {
		*collection = append(*collection, l)
	}
}
