package imageflux

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-ini/ini"
)

var (
	imagefluxConfPath string
)

func init() {
	imagefluxConfPath = fmt.Sprintf("%s/.imageflux/conf.ini", os.Getenv("HOME"))
}

func loadTokenBytes(dat []byte) (string, error) {
	conf, err := ini.Load(dat)
	if err != nil {
		return "", err
	}

	token := ""
	for _, s := range conf.Sections() {
		if s.Name() == "DEFAULT" {
			token = s.Key("token").String()
		} else {
			continue
		}
	}
	return token, nil
}

func loadToken(confPath string) (string, error) {
	dat, err := ioutil.ReadFile(confPath)
	if err != nil {
		return "", err
	}
	return loadTokenBytes(dat)
}
