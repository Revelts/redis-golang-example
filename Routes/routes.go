package Routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"redis-api/Controllers"
)

func HandleRequests() (err error) {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/redis", func(r chi.Router) {
			r.Post("/createTask", Controllers.CreateTask)
		})
	})
	err = http.ListenAndServe(":3000", r)
	return
}
