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