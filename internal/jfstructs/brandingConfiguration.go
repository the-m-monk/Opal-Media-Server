package jfstructs

/*
GET https://demo.jellyfin.org/stable/Branding/Configuration
{
	"LoginDisclaimer":"(Paragraph about something or other)",
	"CustomCss":"",
	"SplashscreenEnabled":true
}
*/

type ResponseBrandingConfiguration struct {
	LoginDisclaimer     string `json:"LoginDisclaimer"`
	CustomCss           string `json:"CustomCss"`
	SplashscreenEnabled bool   `json:"SplashscreenEnabled"`
}
