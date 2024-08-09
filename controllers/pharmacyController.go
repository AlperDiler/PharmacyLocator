package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	models "module/models"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/jftuga/geodist"
	"gorm.io/gorm"
)

func parseCoordinates(pharmacies []models.Pharmacy) {
	for i := range pharmacies {
		coords := pharmacies[i].Koordinat
		parts := strings.Split(coords, ",")

		if len(parts) != 2 {
			continue
		}
		latStr := strings.TrimSpace(parts[0])
		lonStr := strings.TrimSpace(parts[1])

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			log.Println("Error parsing latitude:", err)
			continue
		}
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			log.Println("Error parsing longitude:", err)
			continue
		}
		pharmacies[i].Latidute = lat
		pharmacies[i].Longidute = lon
	}
}
func FindNearestPharmacies(db *gorm.DB, pharmacies []*models.Pharmacy, user []*models.UserCoords) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			http.Error(w, "Database not initialized", http.StatusInternalServerError)
			return
		}
		var userInfo models.UserCoords
		err2 := json.NewDecoder(r.Body).Decode(&userInfo)
		if err2 != nil {
			http.Error(w, "USERINFO JSON CONVERTION : "+err2.Error(), http.StatusBadRequest)
			return
		}
		parsedLat, err := strconv.ParseFloat(userInfo.UserLat, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
		}
		parsedLng, err := strconv.ParseFloat(userInfo.UserLon, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
		}
		userLocation := geo.Location{Lat: parsedLat, Lng: parsedLng}
		fmt.Printf("User Info || UserLat : %v -- UserLon : %v\n", userInfo.UserLat, userInfo.UserLon)
		if err != nil {
			http.Error(w, "INFO ERROR"+err.Error(), http.StatusBadRequest)
		}
		var pharmacies []models.Pharmacy
		db.Find(&pharmacies)

		type distancePharmacy struct {
			ID        int    `json:"id"`
			EczaneAdi string `json:"eczaneAdi"`
			Telefon   string `json:"telefon"`
			Adres     string `json:"adres"`
			Distance  float64
		}

		parseCoordinates(pharmacies)
		var distances []distancePharmacy
		userCoords := geodist.Coord{Lat: userLocation.Lat, Lon: userLocation.Lng}
		for _, pharmacy := range pharmacies {
			pharmacyCoord := geodist.Coord{Lat: pharmacy.Latidute, Lon: pharmacy.Longidute}
			_, km := geodist.HaversineDistance(userCoords, pharmacyCoord)
			distances = append(distances, distancePharmacy{
				ID:        int(pharmacy.ID),
				EczaneAdi: pharmacy.EczaneAdi,
				Telefon:   pharmacy.Telefon,
				Adres:     pharmacy.Adres,
				Distance:  km,
			})
			log.Printf("Pharmacy %s - distance: %.2f km", pharmacy.EczaneAdi, km)
		}
		sort.Slice(distances, func(i, j int) bool {
			return distances[i].Distance < distances[j].Distance
		})
		if len(distances) > 5 {
			distances := distances[:5]
			fmt.Printf("\nDistances : %v\n", distances)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(distances)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Printf("\nClosest Pharmacies : \n")
			for i, d := range distances {
				fmt.Printf("\n%d-)Pharmacy: %s,\nAdress: %v \nDistance: %.2f km\n", i+1, d.EczaneAdi, d.Adres, d.Distance)
			}
		}
	}
}
