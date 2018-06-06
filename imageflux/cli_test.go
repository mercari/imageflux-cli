package imageflux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignature(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		secret string
		path   string
		sig    string
	}{
		{
			secret: "testsigningsecret",
			path:   "/images/1.jpg",
			sig:    "1.-Yd8m-5pXPihiZdlDATcwkkgjzPIC9gFHmmZ3JMxwS0=",
		},
		{
			secret: "testsigningsecret",
			path:   "/c/w=200/images/1.jpg",
			sig:    "1.tiKX5u2kw6wp9zDgl1tLiOIi8IsoRIBw8fVgVc0yrNg=",
		},
	}

	for _, test := range tests {
		sig := signature(test.path, test.secret)
		assert.Equal(test.sig, sig)
	}
}
