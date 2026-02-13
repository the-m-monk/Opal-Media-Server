package usermgmt

import (
	"fmt"
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
			empty_db, empty_err := os.Create("./db/users")
			if empty_err != nil {
				fmt.Println("Severe warning: failed to create ./db/users, authentication will always fail")
				fmt.Printf("Reason: %s\n", empty_err.Error())
			}
			empty_db.Close()
			return
		}

		fmt.Println("Severe warning: failed to read ./db/users, keeping old UserDB in memory")
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	for _, entry := range strings.Split(string(rawDB), "\n") {
		entry_pieces := strings.Split(entry, ":")

		if len(entry_pieces) != 6 { //invalid entry
			continue
		}

		dbEntry := UserRecord{
			Username:     entry_pieces[0],
			HashAlgo:     entry_pieces[1],
			AlgoVersion:  entry_pieces[2],
			Options:      entry_pieces[3],
			Salt:         entry_pieces[4],
			PasswordHash: entry_pieces[5],
		}

		UserDB = append(UserDB, dbEntry)
	}
}

func WriteUserDB() {
	//TODO: implement
}
