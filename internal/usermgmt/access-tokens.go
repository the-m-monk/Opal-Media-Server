package usermgmt

import "sync"

type AccessToken struct {
	TokenString string
	UserId      string
	UserName    string
}

var CurrentTokens []AccessToken
var CurrentTokensMutex sync.Mutex

//TODO: implement token watchdog
