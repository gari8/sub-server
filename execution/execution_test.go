package execution

import (
	"errors"
	fileManager "github.com/gari8/sub-server/file-manager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestFilePath = "TEST"
)

func getFileManager(t *testing.T) *MockFileManager {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockFileManager(ctrl)
}

func TestNew(t *testing.T) {
	fm := fileManager.FileManager{
		FileName:    TestFilePath,
	}
	assert.Equal(t, Execution{fm}, New(fm))
}

func TestRunInit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fm := getFileManager(t)
		fm.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
		e := New(fm)
		assert.Equal(t, nil, e.RunInit())
	})

	t.Run("error", func(t *testing.T) {
		fm := getFileManager(t)
		err := errors.New("TEST")
		fm.EXPECT().Create(gomock.Any()).Return(err).AnyTimes()
		e := New(fm)
		assert.Equal(t, err, e.RunInit())
	})
}

func TestRunServe(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fm := getFileManager(t)
		fm.EXPECT().Read().Return([]byte(fileContent), nil).AnyTimes()
		e := New(fm)

		err := e.RunServe()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		err := errors.New("TEST")
		fm := getFileManager(t)
		fm.EXPECT().Read().Return(nil, err).AnyTimes()
		e := New(fm)

		err = e.RunServe()
		assert.Error(t, err)
	})
}
