package myfuzz
import "github.com/cloudflare/gokey"

func Fuzz(data []byte) int {
	_, err := gokey.GenerateEncryptedKeySeed(string(data))
	if err != nil { return 1 }
	return 0
}
