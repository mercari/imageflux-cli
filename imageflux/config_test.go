package imageflux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfEmptyTokenError(t *testing.T) {
	assert := assert.New(t)

	tests := [][]byte{
		[]byte(""),
		[]byte("token = "),
		[]byte("token = ''"),
		[]byte("token = \"\""),
	}

	for _, test := range tests {
		token, err := loadTokenBytes(test)
		assert.Equal("", token)
		assert.Nil(err)
	}
}

func TestLoadConfParseError(t *testing.T) {
	assert := assert.New(t)

	tests := [][]byte{
		[]byte("dummy"),
		[]byte("token"),
	}

	for _, test := range tests {
		token, err := loadTokenBytes(test)
		assert.Equal("", token)
		assert.NotNil(err)
	}
}

func TestLoadConfSuccess(t *testing.T) {
	assert := assert.New(t)

	tests := [][]byte{
		[]byte("token = 'fffff'"),
		[]byte("token = \"fffff\""),
		[]byte("token = fffff"),
	}

	for _, test := range tests {
		token, err := loadTokenBytes(test)
		assert.Equal("fffff", token)
		assert.Nil(err)
	}
}
