package utils

import (
	"time"

	"github.com/google/uuid"
)

// CreateID id用のユニーク文字列作成
func CreateID() string {
	return time.Now().Format("20060102150405") + uuid.Must(uuid.NewRandom()).String()
}
