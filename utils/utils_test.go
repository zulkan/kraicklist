package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimLowerString(t *testing.T) {
	title := " No 1 Item"
	assert.Equal(t, "no 1 item", TrimLowerString(title))
}

func TestExtractTitle(t *testing.T) {
	title := "iphone for man"
	assert.Equal(t, "iphone", ExtractTitle(title))

	title = "iphone X for better future"
	assert.Equal(t, "iphone X", ExtractTitle(title))
}
