package util

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func GeneratePassword(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateMD5Token(id string) (string, string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(id), bcrypt.DefaultCost)

	hasher := md5.New()
	hasher.Write(hash)

	token := hex.EncodeToString(hasher.Sum(nil))
	return string(hash), token
}

