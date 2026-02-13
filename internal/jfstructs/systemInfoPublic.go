package jfstructs

/*
GET https://demo.jellyfin.org/stable/System/Info/Public
{
	"LocalAddress":"http://172.17.0.2:8096/stable",
	"ServerName":"Stable Demo",
	"Version":"10.11.6",
	"ProductName":"Jellyfin Server",
	"OperatingSystem":"",
	"Id":"f0b3381645f04afb9a0e392e74b6a1b0",
	"StartupWizardCompleted":true
}
*/

type ResponseSystemInfoPublic struct {
	LocalAddress           string `json:"LocalAddress"`
	ServerName             string `json:"ServerName"`
	Version                string `json:"Version"`
	ProductName            string `json:"ProductName"`
	OperatingSystem        string `json:"OperatingSystem"`
	Id                     string `json:"Id"`
	StartupWizardCompleted bool   `json:"StartupWizardCompleted"`
}
