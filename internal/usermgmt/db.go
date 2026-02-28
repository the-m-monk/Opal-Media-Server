package usermgmt

import (
	"log/slog"
	"os"
	"strings"
)

type UserRecord struct {
	Username     string
	HashAlgo     string
	AlgoVersion  string
	Options      string
	Salt         string
	PasswordHash string
}

// TODO: add mutex
var UserDB []UserRecord

func ReadUserDB() {
	//TODO: make this configurable
	rawDB, err := os.ReadFile("./db/users")
	if err != nil {
		if err.Error() == "open ./db/users: no such file or directory" {
			emptyDB, emptyErr := os.Create("./db/users")
			if emptyErr != nil {
				slog.Error("failed to create ./db/users, authentication will always fail", "reason", emptyErr.Error())
			}
			emptyDB.Close()
			return
		}

		slog.Error("failed to read ./db/users, keeping old UserDB in memory", "reason", err.Error())
		return
	}

	for _, entry := range strings.Split(string(rawDB), "\n") {
		entryPieces := strings.Split(entry, ":")

		if len(entryPieces) != 6 { //invalid entry
			continue
		}

		dbEntry := UserRecord{
			Username:     entryPieces[0],
			HashAlgo:     entryPieces[1],
			AlgoVersion:  entryPieces[2],
			Options:      entryPieces[3],
			Salt:         entryPieces[4],
			PasswordHash: entryPieces[5],
		}

		UserDB = append(UserDB, dbEntry)
	}
}

func WriteUserDB() {
	//TODO: implement
}
