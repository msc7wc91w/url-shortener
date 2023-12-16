package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "123456789"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.google.com/"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.facebook.com/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://www.netflix.com/"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "")
	assert.NotEqual(t, shortLink_2, "")
	assert.NotEqual(t, shortLink_3, "")
}
