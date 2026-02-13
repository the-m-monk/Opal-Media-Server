package usermgmt

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"opal/internal/idgen"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

func Authenticate(user UserRecord, passWord string, aT *AccessToken) string {
	//TODO: replace return string with err type

	if user.HashAlgo != "argon2id" ||
		user.AlgoVersion != "19" {
		fmt.Printf("Usermgmt error: fatal database error, user: %s has a corrupted/bad entry in user database\n", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	//optSplit = [m=memory (KiB), t=timecost, p=threads (the number of threads used to create the hash initially, not important)]
	optSplit := strings.Split(user.Options, ",")

	if len(optSplit) != 3 {
		fmt.Printf("Usermgmt error: fatal database error, user: %s has a corrupted/bad entry in user database\n", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	mem, err := strconv.Atoi(strings.Split(optSplit[0], "=")[1])
	if err != nil {
		fmt.Printf("Usermgmt error while parsing %s's entry: %s\n", user.Username, err.Error())
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	if mem < (19 * 1024) {
		//TODO: multiline this print statement, it's a bit long
		fmt.Printf("SEVERE WARNING: The recommended minimum amount of memory for use with argon2id is 19Mib, %s's configured argon2id authentication memory is only %dKiB. \nThis poses a potential security risk and should be updated immediately to a higher value\n", user.Username, mem)
	}

	timeCost, err := strconv.Atoi(strings.Split(optSplit[1], "=")[1])
	if err != nil {
		fmt.Printf("Usermgmt error while parsing %s's entry: %s\n", user.Username, err.Error())
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	if timeCost < 2 {
		//TODO: multiline this print statement, it's a bit long
		fmt.Printf("SEVERE WARNING: The minimum recommended timecost for use with argon2id is 2, %s's configured argon2id timecost is only %d. \nThis poses a potential security risk and should be updated immediately to a higher value\n", user.Username, timeCost)
	}

	salt, err := base64.RawStdEncoding.DecodeString(user.Salt)
	if err != nil {
		fmt.Printf("Usermgmt: failed to decode salt: %s\n", err.Error())
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	//TODO: change timecost and memory to be explicity uint32 when declared
	attemptedPhash := argon2.IDKey([]byte(passWord), salt, uint32(timeCost), uint32(mem), 1, 32)

	phash, err := base64.RawStdEncoding.DecodeString(user.PasswordHash)
	if err != nil {
		//TODO: increase the verbosity of this print (I got lazy)
		fmt.Println("Usermgmt: failed to decode password hash, potential database corruption")
	}

	if !bytes.Equal(phash, attemptedPhash) {
		return "FAILED_TO_AUTHENTICATE"
	}

	tokenString, err := idgen.GenerateRandomId(32)
	if err != nil {
		//TODO: increase the verbosity of this print (I got lazy)
		fmt.Println("Usermgmt: failed to generate token string")
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	aT.TokenString = tokenString
	return "SUCCESS"
}
