package file_manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testFp = "test/index.go"
)

func TestNewFileManager(t *testing.T) {
	assert.Equal(t, FileManager{FileName: testFp}, NewFileManager(testFp))
}

func TestFileManager_Create(t *testing.T) {
	t.Skip()
	// TODO
}