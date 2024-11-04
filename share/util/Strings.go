package util

import (
	"math/rand"
	"time"
)

func GenerateRandomStringNumber(length int) string {
    rand.Seed(time.Now().UnixNano()) 

    // 使用する文字のリスト
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    // ランダムな文字列を生成
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}