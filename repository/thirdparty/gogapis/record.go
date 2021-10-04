package gogapis

import (
	"fmt"
	"gameprice-api/business/gogapis"
)

type GOG struct {
	ID                         int    `json:"id"`
	Title                      string `json:"title"`
	PurchaseLink               string `json:"purchase_link"`
	Slug                       string `json:"slug"`
	ContentSystemCompatibility struct {
		Windows bool `json:"windows"`
		Osx     bool `json:"osx"`
		Linux   bool `json:"linux"`
	} `json:"content_system_compatibility"`
	Languages struct {
		Cn string `json:"cn"`
		De string `json:"de"`
		En string `json:"en"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		Jp string `json:"jp"`
		Ko string `json:"ko"`
		Pt string `json:"pt"`
		Ru string `json:"ru"`
	} `json:"languages"`
	Links struct {
		PurchaseLink string `json:"purchase_link"`
		ProductCard  string `json:"product_card"`
		Support      string `json:"support"`
		Forum        string `json:"forum"`
	} `json:"links"`
	InDevelopment struct {
		Active bool        `json:"active"`
		Until  interface{} `json:"until"`
	} `json:"in_development"`
	IsSecret      bool   `json:"is_secret"`
	IsInstallable bool   `json:"is_installable"`
	GameType      string `json:"game_type"`
	IsPreOrder    bool   `json:"is_pre_order"`
	ReleaseDate   string `json:"release_date"`
	Images        struct {
		Background          string `json:"background"`
		Logo                string `json:"logo"`
		Logo2X              string `json:"logo2x"`
		Icon                string `json:"icon"`
		SidebarIcon         string `json:"sidebarIcon"`
		SidebarIcon2X       string `json:"sidebarIcon2x"`
		MenuNotificationAv  string `json:"menuNotificationAv"`
		MenuNotificationAv2 string `json:"menuNotificationAv2"`
	} `json:"images"`
	Dlcs []interface{} `json:"dlcs"`
}

func (rec *GOG) ToDomain() gogapis.Domain {
	fmt.Printf("%+v", rec)
	return gogapis.Domain{
		URL: rec.Links.ProductCard,
	}
}