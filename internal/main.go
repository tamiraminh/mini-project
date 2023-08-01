package internal

import (
	"mini-project/docs"
	"mini-project/internal/handler"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)


// @title Booking API Docummentaion
// @description This is a sample API for managing bookings.
// @version 1.0
// @host localhost:8080
// @BasePath /
func RunServer() {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	})
	
	docs.SwaggerInfo.Title = "Booking Photoshoot Management"
    docs.SwaggerInfo.Version = "v1"
    conf := httpSwagger.URL("http://localhost:8080/swagger/doc.json")
	
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsOptions.Handler)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	
	
	r.Route("/v1/bookings", func(r chi.Router) {
		r.Get("/", handler.ReadBookings) 
        r.Post("/", handler.CreateBooking)
        r.Put("/{id}" , handler.UpdateBooking)
        r.Delete("/{id}", handler.DeleteBooking)
    })

	r.Get("/swagger/*", httpSwagger.Handler(conf))
	
	http.ListenAndServe(":8080", r)
}
