package entity

import "time"

type SkinPrice struct {
	Name            string  `json:"name"`
	Float           float64 `json:"float"`
	QualityValue    string  `json:"qualityValue"`
	MarketplaceName string  `json:"marketplaceName"`
	MarketplaceUrl  string  `json:"marketplaceUrl"`
	USDPrice        int64   `json:"USDPrice"`
}

type WebSite struct {
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	TimeVisited time.Time `json:"timeVisited"`
	Browser     string    `json:"browser"`
	OS          string    `json:"os"`
	Region      string    `json:"region"`
}

func GetSiteHeaders() string {
	return "name" + "," + "url" + "," + "timeVisited" + "," + "browser" + "," + "os" + "," + "region"
}

func (site *WebSite) GetValues() string {
	return site.Name + "," + site.Url + "," + site.TimeVisited.String() + "," + site.Browser + "," + site.OS + "," + site.Region
}
