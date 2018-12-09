package staticModel

type configContainer struct {
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
	domain        string
	disqusId      string
	targetDir     string
	description   string
	homeText      string
	homeHeadline  string
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

func (cc *configContainer) DisqusId() string { return cc.disqusId }

func (cc *configContainer) TargetDir() string { return cc.targetDir }

func (cc *configContainer) HomeText() string { return cc.homeText }

func (cc *configContainer) HomeHeadline() string { return cc.homeHeadline }
