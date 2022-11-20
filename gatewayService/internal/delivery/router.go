package delivery

import (
	"github.com/mohammaderm/todoMicroService/gatewayService/config"

	"net/http"

	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/monitoring"

	_ "github.com/mohammaderm/todoMicroService/gatewayService/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RouteProvider struct {
	AuthHandler     AuthHandlerContract
	TodoHandler     TodoHandlerContrac
	CategoryHandler CategoryHandlerContrac
	Monitoring      monitoring.MetricsCallectors
	Cfg             *config.Token
}

func RouterProvider(rp *RouteProvider) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(MetricMiddleware(rp.Monitoring))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(cors.Handler)

	r.Route("/todo", func(r chi.Router) {
		r.Use(cors.Handler)
		r.Use(Auth(rp.Cfg))
		r.Use(MetricMiddleware(rp.Monitoring))
		r.Post("/create", rp.TodoHandler.Create)
		r.Get("/getAll", rp.TodoHandler.GetAll)
		r.Delete("/{id}", rp.TodoHandler.Delete)
		r.Put("/{id}", rp.TodoHandler.Update)
	})

	r.Route("/category", func(r chi.Router) {
		r.Use(cors.Handler)
		r.Use(Auth(rp.Cfg))
		r.Use(MetricMiddleware(rp.Monitoring))
		r.Post("/create", rp.CategoryHandler.CreateCat)
		r.Get("/getAll", rp.CategoryHandler.GetAllCat)
		r.Delete("/{id}", rp.CategoryHandler.DeleteCat)
	})

	r.Post("/auth/login", rp.AuthHandler.Login)
	r.Post("/auth/register", rp.AuthHandler.Register)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	// r.Mount("/api/v1", r)

	return r
}
