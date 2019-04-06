package staticModel

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticPersistence"
)

func NewPageMaker() staticIntf.PageMaker { return &pageMaker{} }

type pageMaker struct {
	title           string
	description     string
	content         string
	category        string
	createDate      string
	pathFromDocRoot string
	fileName        string
	tags            []string
	images          []staticIntf.Image
	site            staticIntf.Site
	navigatedPages  []staticIntf.Page
}

func (pm *pageMaker) Title(title string)                     { pm.title = title }
func (pm *pageMaker) Description(description string)         { pm.description = description }
func (pm *pageMaker) Content(content string)                 { pm.content = content }
func (pm *pageMaker) Category(category string)               { pm.category = category }
func (pm *pageMaker) CreateDate(createDate string)           { pm.createDate = createDate }
func (pm *pageMaker) PathFromDocRoot(pathFromDocRoot string) { pm.pathFromDocRoot = pathFromDocRoot }
func (pm *pageMaker) FileName(fileName string)               { pm.fileName = fileName }
func (pm *pageMaker) Tags(tags ...string)                    { pm.tags = tags }
func (pm *pageMaker) Images(images ...staticIntf.Image)      { pm.images = images }
func (pm *pageMaker) Site(site staticIntf.Site)              { pm.site = site }
func (pm *pageMaker) NavigatedPages(navigatedPages ...staticIntf.Page) {
	pm.navigatedPages = navigatedPages
}

func (pm *pageMaker) Make() staticIntf.Page {
	dto := staticPersistence.NewPageDto(
		pm.title,
		pm.description,
		pm.content,
		pm.category,
		pm.createDate,
		pm.pathFromDocRoot,
		pm.fileName,
		pm.tags,
		pm.images)
	newPage := NewPage(
		dto,
		pm.site)
	newPage.NavigatedPages(pm.navigatedPages...)
	return newPage
}
