package entity

import (
	"math/rand"
	"strings"
	"time"
)

const (
	BitSkins   = "BitSkins"
	DMarket    = "DMarket"
	CSMoney    = "CSMoney"
	Tradeit    = "TradeIT GG"
	CSGOTM     = "CSGO TM"
	SkinWallet = "SkinWallet"
	Yandex     = "Yandex"
	Google     = "Google"
	VK         = "VK"
	YouTube    = "YouTube"
	MailRu     = "MailRu"
)

const (
	BitSkinsURL   = "bitskins.com"
	DMarktetURL   = "dmarket.ru"
	CSMoneyURL    = "csmoney.com"
	TradeitURL    = "tradeit.gg"
	SkinWalletURL = "skinwallet.com"
	CSGOTMURL     = "csgotm.ru"
	YandexURL     = "ya.ru"
	GoogleURL     = "google.com"
	VKURL         = "vk.ru"
	YouTubeURL    = "youtube.com"
	MailRuURL     = "mail.ru"
)

const (
	Ak47Redline    = "AK - 47 | Redline"
	AwpAsiimov     = "AWP | Asiimov"
	AwpDragonLore  = "AWP | Dragon Lore"
	PpBison        = "PP-19 | Bison"
	FamasPule      = "Famas | Pulse"
	Ak47Vulkan     = "AK-47 | Vulkan"
	M4a1sdecimator = "M4A1-S | Decimator"
)

const (
	Chrome = "Google Chrome"
	Opera  = "Opera"
	MSEdge = "Microsoft Edge"
	Safari = "Apple Safari"
	Ishak  = "Internet Explorer"
)

const (
	Win10   = "Windows 10"
	Win8    = "Windows 8"
	Win7    = "Windows 7"
	MacOS   = "macOS"
	Ubuntu  = "Ubuntu Linux"
	Debian  = "Debian Linux"
	Ios     = "Apple iOS"
	Android = "Android"
)

type WebSiteGenerator struct {
	sites    map[string]string
	browsers []string
	oses     []string
}

func NewSiteGen() (error, *WebSiteGenerator) {

	sites := make(map[string]string)

	sites[BitSkins] = BitSkinsURL
	sites[DMarket] = DMarktetURL
	sites[CSMoney] = CSMoneyURL
	sites[Tradeit] = TradeitURL
	sites[SkinWallet] = SkinWalletURL
	sites[CSGOTM] = CSGOTMURL
	sites[Yandex] = YandexURL
	sites[Google] = GoogleURL
	sites[VK] = VKURL
	sites[YouTube] = YouTubeURL
	sites[MailRu] = MailRuURL

	browsers := make([]string, 0)

	browsers = append(browsers,
		Chrome, MSEdge, Opera, Safari, Ishak,
	)

	oses := make([]string, 0)

	oses = append(oses,
		Win7, Win8, Win10, MacOS, Ubuntu, Debian, Ios, Android,
	)

	sitesGen := &WebSiteGenerator{
		sites:    sites,
		browsers: browsers,
		oses:     oses,
	}

	return nil, sitesGen
}

func (sg *WebSiteGenerator) GenerateSite() (error, WebSite) {
	siteKeys := getKeys(sg.sites)

	site := siteKeys[rand.Intn(len(siteKeys))]
	siteUrl := sg.sites[site]

	siteEntity := WebSite{
		Name:        site,
		Url:         siteUrl,
		OS:          sg.oses[rand.Intn(len(sg.oses))],
		Browser:     sg.browsers[rand.Intn(len(sg.browsers))],
		Region:      domainType(siteUrl),
		TimeVisited: randomTimeFromYearStart(),
	}
	return nil, siteEntity
}

type SkinPriceGenerator struct {
	marketplaces map[string]string
	skins        []string
}

func NewGenerator() (error, *SkinPriceGenerator) {

	marketplaces := make(map[string]string)

	marketplaces[BitSkins] = BitSkinsURL
	marketplaces[DMarket] = DMarktetURL
	marketplaces[CSMoney] = CSMoneyURL
	marketplaces[Tradeit] = TradeitURL
	marketplaces[SkinWallet] = SkinWalletURL

	skins := make([]string, 0)

	skins = append(skins,
		Ak47Redline, AwpAsiimov, AwpDragonLore, PpBison, FamasPule, Ak47Vulkan, M4a1sdecimator,
	)

	generator := &SkinPriceGenerator{
		marketplaces: marketplaces,
		skins:        skins,
	}

	return nil, generator
}

func (gen *SkinPriceGenerator) GenerateSkinPrice() (error, SkinPrice) {
	skinFloat := rand.Float64()

	var quality string

	switch true {
	case skinFloat > 0.45:
		quality = "Battle-Scarred"
	case skinFloat > 0.38 && skinFloat <= 0.45:
		quality = "Well-Worn"
	case skinFloat > 0.15 && skinFloat <= 0.38:
		quality = "Field Tested"
	case skinFloat > 0.7 && skinFloat <= 0.15:
		quality = "Minimal Wear"
	case skinFloat >= 0 && skinFloat <= 0.7:
		quality = "Factory New"

	default:
		quality = "None"
	}

	keys := getKeys(gen.marketplaces)

	marketplaceName := keys[rand.Intn(len(keys))]

	marketplaceUrl := gen.marketplaces[marketplaceName]

	randInt := rand.Intn(len(gen.skins))

	sp := SkinPrice{
		Name:            gen.skins[randInt],
		Float:           skinFloat,
		QualityValue:    quality,
		MarketplaceName: marketplaceName,
		MarketplaceUrl:  marketplaceUrl,
		USDPrice:        rand.Int63n(99999) + 1,
	}

	return nil, sp
}

func getKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)

	for i := range m {
		keys = append(keys, i)
	}

	return keys
}

func domainType(domain string) string {
	if strings.HasSuffix(domain, ".ru") {
		return "Russia"
	}
	return "International"
}

func randomTimeFromYearStart() time.Time {
	yearStart := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	randomSeconds := rand.Int63n(time.Now().Unix() - yearStart.Unix())
	randomTime := yearStart.Add(time.Duration(randomSeconds) * time.Second)
	return randomTime
}
