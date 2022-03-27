package services

import (
	"github.com/hiroyky/famiphoto/usecases"
	"math/rand"
	"time"
)

func NewRandomService() usecases.RandomService {
	return &randomService{
		letters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
}

type randomService struct {
	letters string
}

func (s *randomService) GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = s.letters[int(rand.Int63()%int64(len(s.letters)))]
	}
	return string(b)
}
