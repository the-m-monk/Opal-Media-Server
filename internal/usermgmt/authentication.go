package usermgmt

import (
	"bytes"
	"encoding/base64"
	"log/slog"
	"opal/internal/idgen"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

//TODO: Clean all of this code, return err instead of string

func Authenticate(user UserRecord, passWord string, aT *AccessToken) string {
	//TODO: replace return string with err type

	if user.HashAlgo != "argon2id" ||
		user.AlgoVersion != "19" {
		slog.Error("user has a corrupted/bad entry in user database", "user", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	//optSplit = [m=memory (KiB), t=timecost, p=threads (the number of threads used to create the hash initially, not important)]
	optSplit := strings.Split(user.Options, ",")

	if len(optSplit) != 3 {
		slog.Error("user has a corrupted/bad entry in user database", "user", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	mem, err := strconv.Atoi(strings.Split(optSplit[0], "=")[1])
	if err != nil {
		slog.Error("failed to parse user entry", "user", user.Username, "reason", err.Error())
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	if mem < (19 * 1024) {
		slog.Warn("user's authentication memory is less than recommended minimum (19 Mib)", "user", user.Username, "usersCurrentAuthMemInKiB", mem)
	}

	timeCost, err := strconv.Atoi(strings.Split(optSplit[1], "=")[1])
	if err != nil {
		slog.Error("user has a corrupted/bad entry in user database", "user", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	if timeCost < 2 {
		slog.Warn("user's authentication timecost is less than recommended minimum (2)", "user", user.Username, "usersCurrentAuthTimeCost", timeCost)
	}

	salt, err := base64.RawStdEncoding.DecodeString(user.Salt)
	if err != nil {
		slog.Error("user has a corrupted/bad entry in user database", "user", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	//TODO: change timecost and memory to be explicity uint32 when declared
	attemptedPhash := argon2.IDKey([]byte(passWord), salt, uint32(timeCost), uint32(mem), 1, 32)

	phash, err := base64.RawStdEncoding.DecodeString(user.PasswordHash)
	if err != nil {
		slog.Error("user has a corrupted/bad entry in user database", "user", user.Username)
	}

	if !bytes.Equal(phash, attemptedPhash) {
		slog.Warn("user failed to authenticate", "user", user.Username)
		return "FAILED_TO_AUTHENTICATE"
	}

	tokenString, err := idgen.GenerateRandomId(32)
	if err != nil {
		slog.Error("failed to generate token string", "user", user.Username)
		return "INTERNAL_AUTHENTICATION_ERROR"
	}

	aT.TokenString = tokenString
	slog.Info("user authenticated successfully", "user", user.Username)
	return "SUCCESS"
}
