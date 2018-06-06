package imageflux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateError(t *testing.T) {
	assert := assert.New(t)

	tests := []APIParam{
		{ // empty token
			Token: "",
			URL:   "https://example.com/images/example.jpg",
		},
		{ // URL is empty
			Token: "dummy",
			URL:   "",
		},
		{ // key is not URL
			Token: "dummy",
			URL:   "dummy-url",
		},
	}

	for _, test := range tests {
		err := validate(&test)
		assert.NotNil(err)
	}

}

func TestValidateSuccess(t *testing.T) {
	assert := assert.New(t)

	tests := []APIParam{
		{ // http
			Token: "dummy",
			URL:   "http://example.com/images/example.jpg",
		},
		{ // https
			Token: "dummy",
			URL:   "https://example.com/images/example.jpg",
		},
		{ // paramenter required urlencode is involved
			Token: "dummy",
			URL:   "https://example.com/c!/w=200,h=200/images/example.jpg",
		},
	}

	for _, test := range tests {
		err := validate(&test)
		assert.Nil(err)
	}

}
