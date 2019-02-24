package staticModel

type configContainer struct {
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

func (cc *configContainer) TwitterHandle() string { return cc.twitterHandle }

func (cc *configContainer) Description() string { return cc.description }

func (cc *configContainer) Topic() string { return cc.topic }

func (cc *configContainer) Tags() string { return cc.tags }

func (cc *configContainer) CardType() string { return cc.cardType }

func (cc *configContainer) Section() string { return cc.section }

func (cc *configContainer) FBPage() string { return cc.fbPage }

func (cc *configContainer) TwitterPage() string { return cc.twitterPage }

func (cc *configContainer) RssPath() string { return cc.rssPath }

func (cc *configContainer) RssFilename() string { return cc.rssFilename }

func (cc *configContainer) Css() string { return cc.css }

func (cc *configContainer) Domain() string { return cc.domain }

func (cc *configContainer) SvgLogo() string { return cc.svgLogo }

func (cc *configContainer) BasePath() string { return cc.basePath }

func (cc *configContainer) DisqusId() string { return cc.disqusId }

func (cc *configContainer) TargetDir() string { return cc.targetDir }

func (cc *configContainer) HomeText() string { return cc.homeText }

func (cc *configContainer) HomeHeadline() string { return cc.homeHeadline }

func (cc *configContainer) KeyWords() string { return cc.keyWords }

func (cc *configContainer) Subject() string { return cc.subject }

func (cc *configContainer) Author() string { return cc.author }
