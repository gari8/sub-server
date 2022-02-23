package server

import (
	"github.com/gari8/sub-server/setting"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testContent = `{
  "id": 1,
  "text": "content"
}`
)

func getMockOrigin(method string) setting.Origin {
	return setting.Origin{
		URI:      "/app/v1/start",
		FilePath: ".",
		Method:   method,
		Content:  []byte(testContent),
	}
}

func TestNewHttpServer(t *testing.T) {
	assert.Equal(t, HttpServer{}, NewHttpServer())
}

func TestHttpServer_multiplex(t *testing.T) {
	router := chi.NewRouter()
	mockOrigin := getMockOrigin("GET")
	err := multiplex(router, mockOrigin)
	assert.NoError(t, err)
}
