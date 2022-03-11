package services

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func NewPasswordService() usecases.PasswordService {
	return &passwordService{hmacKey: []byte(config.Env.HMacKey)}
}

type passwordService struct {
	hmacKey []byte
}

func (s *passwordService) HashPassword(password string) (string, error) {
	dst, err := bcrypt.GenerateFromPassword(s.hmacHash(password), 30)
	if err != nil {
		return "", errors.New(errors.HashPasswordFatal, err)
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

func (s *passwordService) MatchPassword(password string, correctPassword string) (bool, error) {
	decodedCorrect, err := base64.StdEncoding.DecodeString(correctPassword)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(s.hmacHash(password), decodedCorrect)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, errors.New(errors.MatchPasswordFatal, err)
}

func (s *passwordService) GeneratePassword() (string, error) {
	return password.Generate(20, 10, 0, false, false)
}

func (s *passwordService) hmacHash(src string) []byte {
	h := hmac.New(sha512.New, s.hmacKey)
	h.Write([]byte(src))
	return h.Sum(nil)
}
