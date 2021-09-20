package server

import (
	"fmt"
	"github.com/gari8/sub-server/setting"
	"github.com/gari8/sub-server/tools/format"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type HttpServer struct {}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) Serve(stg setting.Setting) error {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Route(stg.OriginRoot, func(r chi.Router) {
		for _, origin := range stg.Origins {
			fmt.Println(origin.FilePath)
			multiplex(r, origin)
		}
	})

	format.Logo()
	for _, origin := range stg.Origins {
		format.Print(format.PBlue, "\t"+origin.Method+" >> "+"http://localhost:"+stg.Port+stg.OriginRoot+origin.URI)
	}
	return http.ListenAndServe(":"+stg.Port, router)
}

func multiplex(r chi.Router, origin setting.Origin) {
	switch origin.Method {
	case "GET":
		fmt.Println(string(origin.Content))
		r.Get(origin.URI, func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(origin.Content)
		})
	case "POST":
		r.Post(origin.URI, func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(origin.Content)
		})
	case "PUT":
		r.Put(origin.URI, func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(origin.Content)
		})
	case "PATCH":
		r.Patch(origin.URI, func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(origin.Content)
		})
	case "DELETE":
		r.Delete(origin.URI, func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(origin.Content)
		})
	}
}