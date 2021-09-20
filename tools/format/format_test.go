package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlashAssign(t *testing.T) {
	textWithSlash := "/app/v1/target"
	textWithoutSlash := "app/v1/target"
	assert.Equal(t, textWithSlash, SlashAssign(textWithoutSlash))
}

func TestRemoveDotSlash(t *testing.T) {
	textWithDotSlash := "./app/v1/target"
	textWithoutDotSlash := "app/v1/target"
	assert.Equal(t, RemoveDotSlash(textWithDotSlash), textWithoutDotSlash)
}
