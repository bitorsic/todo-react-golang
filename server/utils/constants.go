package utils

import "time"

const (
	AuthTokenExp    = time.Minute * 10
	RefreshTokenExp = time.Hour * 24 * 30 // 1 month
)
