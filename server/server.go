package server

import (
	"github.com/gari8/sub-server/setting"
	"github.com/gari8/sub-server/tools/format"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type HttpServer struct{}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) Serve(stg setting.Setting) (Err error) {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Route(format.SlashAssign(stg.OriginRoot), func(r chi.Router) {
		for _, origin := range stg.Origins {
			err := multiplex(r, origin)
			if err != nil {
				Err = err
				return
			}
		}
	})

	format.Logo()
	for _, origin := range stg.Origins {
		format.Print(format.PBlue, "\t"+origin.Method+" >> "+"http://localhost:"+stg.Port+format.SlashAssign(stg.OriginRoot)+format.SlashAssign(origin.URI))
	}
	return http.ListenAndServe(":"+stg.Port, router)
}

func multiplex(r chi.Router, origin setting.Origin) (Err error) {
	switch origin.Method {
	case "GET":
		r.Get(format.SlashAssign(origin.URI), func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write(origin.Content)
			if err != nil {
				Err = err
				return
			}
		})
	case "POST":
		r.Post(format.SlashAssign(origin.URI), func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write(origin.Content)
			if err != nil {
				Err = err
				return
			}
		})
	case "PUT":
		r.Put(format.SlashAssign(origin.URI), func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write(origin.Content)
			if err != nil {
				Err = err
				return
			}
		})
	case "PATCH":
		r.Patch(format.SlashAssign(origin.URI), func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write(origin.Content)
			if err != nil {
				Err = err
				return
			}
		})
	case "DELETE":
		r.Delete(format.SlashAssign(origin.URI), func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write(origin.Content)
			if err != nil {
				Err = err
				return
			}
		})
	}
	return nil
}
