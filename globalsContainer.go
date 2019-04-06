package staticModel

type globalsContainer struct {
	domain        string
	basePath      string
	twitterHandle string
	topic         string
	tags          string
	cardType      string
	section       string
	fbPage        string
	twitterPage   string
	rssPath       string
	rssFilename   string
	css           string
	disqusId      string
	targetDir     string
	description   string
	homeText      string
	homeHeadline  string
	svgLogo       string
	keyWords      string
	subject       string
	author        string
}

func (cc *globalsContainer) TwitterHandle() string { return cc.twitterHandle }

func (cc *globalsContainer) Description() string { return cc.description }

func (cc *globalsContainer) Topic() string { return cc.topic }

func (cc *globalsContainer) Tags() string { return cc.tags }

func (cc *globalsContainer) CardType() string { return cc.cardType }

func (cc *globalsContainer) Section() string { return cc.section }

func (cc *globalsContainer) FBPage() string { return cc.fbPage }

func (cc *globalsContainer) TwitterPage() string { return cc.twitterPage }

func (cc *globalsContainer) RssPath() string { return cc.rssPath }

func (cc *globalsContainer) RssFilename() string { return cc.rssFilename }

func (cc *globalsContainer) Css() string { return cc.css }

func (cc *globalsContainer) Domain() string { return cc.domain }

func (cc *globalsContainer) SvgLogo() string { return cc.svgLogo }

func (cc *globalsContainer) BasePath() string { return cc.basePath }

func (cc *globalsContainer) DisqusId() string { return cc.disqusId }

func (cc *globalsContainer) TargetDir() string { return cc.targetDir }

func (cc *globalsContainer) HomeText() string { return cc.homeText }

func (cc *globalsContainer) HomeHeadline() string { return cc.homeHeadline }

func (cc *globalsContainer) KeyWords() string { return cc.keyWords }

func (cc *globalsContainer) Subject() string { return cc.subject }

func (cc *globalsContainer) Author() string { return cc.author }
