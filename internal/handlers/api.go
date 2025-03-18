package handlers

import (
  "github.com/go-chi/chi";
  chimiddle "github.com/go-chi/chi/middleware";
  "github.com/sd191100/goapi.git/internal/middleware";
)

func Handler(r *chi.Mux) {

  // for striping any trailing slashes eg: https://localhost:8000/api/hello/ the last slash will be emmited
  r.Use(chimiddle.StripSlashes); 
  r.Route("/account", func(router chi.Router) {
    router.Use(middleware.Authorization);

    router.Get("/coins", GetCoinBalance);
    
  });
}
