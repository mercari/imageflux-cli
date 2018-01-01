package imageflux

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

const (
	signedURLVersion = 1
)

func buildHMACSHA256(msg, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

func signature(path, secret string) string {
	return fmt.Sprintf("%d.%s", signedURLVersion, buildHMACSHA256(path, secret))
}

func Signature(command string) (string, error) {
	f := flag.NewFlagSet(os.Args[0]+" "+command, flag.ExitOnError)
	secret := f.String("s", "", "signing secret")
	path := f.String("p", "", "url apart from scheme and hostname")
	f.Parse(os.Args[2:])

	if *secret == "" {
		return "", fmt.Errorf("secret is empty")
	}

	if *path == "" {
		return "", fmt.Errorf("path is empty")
	}

	return signature(*path, *secret), nil
}
