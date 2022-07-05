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
			r.Post("/", Controllers.RedisSetter)
			r.Get("/{key}", Controllers.RedisGetter)
			r.Get("/getTask/{username}", Controllers.GetTask)
			r.Get("/getCompanyTask/{username}", Controllers.GetCompanyTask)
			r.Post("/createTask", Controllers.CreateTask)
			r.Delete("/deleteAllUserTask/{username}", Controllers.DeleteAllUserTask)
			r.Delete("/deleteChoosenCompanyTaskFromUser/{username}", Controllers.DeleteChoosenCompanyTaskFromUser)
		})
	})
	err = http.ListenAndServe(":3000", r)
	return
}
