package utils

import (
	"github.com/deatil/go-encoding/base62"
	"github.com/google/uuid"
)

func GenerateShortCode() string {
	u := uuid.New()
	return string(base62.StdEncoding.Encode([]byte(u.String()[:6])))
}
