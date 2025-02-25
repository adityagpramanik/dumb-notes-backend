package routes

import (
	"dumbnotes.com/dumb-notes-backend/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // MARK: Root routes
    router.HandleFunc("/", controllers.RootHandler).Methods("GET")

    return router
}
