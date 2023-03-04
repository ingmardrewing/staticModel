package staticModel

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// pageContent
type pageContent struct {
	doc           *htmlDoc.HtmlDoc
	id            int
	content       string
	description   string
	imageUrl      string
	publishedTime string
	disqusId      string
	thumbBase64   string
	category      string
	images        []staticIntf.Image
	createdDate   string
}

func (p pageContent) GetDoc() *htmlDoc.HtmlDoc { return p.doc }

func (p pageContent) Category() string { return p.category }

func (p pageContent) Id() int { return p.id }

func (p pageContent) ThumbBase64() string { return p.thumbBase64 }

func (p pageContent) DisqusId() string { return p.disqusId }

func (p pageContent) Content() string { return p.content }

func (p pageContent) BodyText() string {
	// ToDo: Fix this in the dto and persisted data
	// we need the text in a separate field
	parts := strings.Split(p.content, "</a>")
	partsLessFirstLinkedImage := parts[1:len(parts)]
	contentWithoutImage := strings.Join(partsLessFirstLinkedImage, "</a>")
	return contentWithoutImage
}

func (p pageContent) Description() string { return p.description }

func (p pageContent) Images() []staticIntf.Image { return p.images }

func (p pageContent) ImageUrl() string { return p.imageUrl }

func (p pageContent) CreatedDate() string { return p.createdDate }

func (p pageContent) PublishedTime(format ...string) string {
	rx := regexp.MustCompile("(\\d{4})-(\\d{1,2})-(\\d{1,2}) (\\d{1,2}):(\\d{1,2}):(\\d{1,2})")
	m := rx.FindStringSubmatch(p.publishedTime)

	if len(m) > 1 {
		m := rx.FindStringSubmatch(p.publishedTime)
		conv := func(a string) int { i, _ := strconv.Atoi(a); return i }
		loc, _ := time.LoadLocation("Europe/Berlin")
		t := time.Date(
			conv(m[1]),
			time.Month(conv(m[2])),
			conv(m[3]),
			conv(m[4]),
			conv(m[5]),
			conv(m[6]),
			0,
			loc)
		if len(format) > 0 {
			return t.Format(format[0])
		}
		stamp := t.Format("Mon, 02 Jan 2006 15:04:05")
		return stamp + " +0100"
	}

	return p.publishedTime
}
