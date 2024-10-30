package server

import (
	"github.com/assaidy/url-shortener/handlers"
)

func (s *FiberServer) RegisterRoutes() {
	urlH := handlers.NewURLHandler(s.DB)

	s.Post("/shorten", urlH.HandleCreateURL)
	s.Get("/shorten/:sc", urlH.HandleGetURL) // sc = short code
	s.Put("/shorten/:sc", urlH.HandleUpdateURL)
	s.Delete("/shorten/:sc", urlH.HandleDeleteURL)
	s.Get("/shorten/:sc/stats", urlH.HandleGetURLWithStats)
}
