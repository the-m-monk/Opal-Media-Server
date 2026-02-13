package usermgmt

import (
	"github.com/google/uuid"
)

// random uuid 5 namespace:
const uuidNamespace = "1addebfc-f8e8-11f0-823b-325096b39f47"

func GetId(username string) string {
	ns := uuid.MustParse(uuidNamespace)
	u := uuid.NewSHA1(ns, []byte(username))

	return u.String()
}
