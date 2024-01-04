package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" // package for web development
	"github.com/ishantiw/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service....")

	fmt.Println(`
                                                       
	_____        _____ _____ _____    ____                
   |   __|___   |  _  |  _  |     |  |    \ ___ _____ ___ 
   |  |  | . |  |     |   __|-   -|  |  |  | -_|     | . |
   |_____|___|  |__|__|__|  |_____|  |____/|___|_|_|_|___|
														  
   `)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
