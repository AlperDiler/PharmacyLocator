package routers

import (
	"module/controllers"
	"module/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/pharmacies", controllers.FindNearestPharmacies(db, []*models.Pharmacy{}, []*models.UserCoords{})).Methods("POST")
	http.ListenAndServe(":3000", router)
}
