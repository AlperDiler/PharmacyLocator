package main

import (
	"log"
	"module/models"
	"module/routers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./pharmacy.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Pharmacy{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table Created .")

	/* 	eczaneler := []models.Pharmacy{
	   		{EczaneAdi: "GAYRET ECZANESİ", Telefon: "0236 768 14 78", Adres: "ALTEYLÜL MAHALLESİ HÜKÜMET CADDESİ NO:1 AHMETLİ", Ilce: "AHMETLİ", Tarih: "2020-08-26", Koordinat: "38.519300,27.938400"},
	   		{EczaneAdi: "OZAN ECZANESİ", Telefon: "02364144730", Adres: "PAŞA MAHALLESİ 19.SOKAK NO:73 AKHİSAR", Ilce: "AKHİSAR", Tarih: "2020-08-26", Koordinat: "38.920600,27.837600"},
	   		{EczaneAdi: "ÖRMECİ ECZANESİ", Telefon: "0236 413 28 56", Adres: "PAŞA MAH. 19 SOK. NO:81/B AKHİSAR", Ilce: "AKHİSAR", Tarih: "2020-08-26", Koordinat: "38.920600,27.837900"},
	   		{EczaneAdi: "DOĞAN ECZANESİ", Telefon: "0236 654 30 00", Adres: "YENİCE MAHALLESİ KEBAN SOKAK NO:12 ALAŞEHİR", Ilce: "ALAŞEHİR", Tarih: "2020-08-26", Koordinat: "38.350900,28.512900"},
	   		{EczaneAdi: "EGE ECZANESİ", Telefon: "02364624017", Adres: "CUMHURİYET MAHALLESİ MENDERES BULVARI NO:6 DEMİRCİ", Ilce: "DEMİRCİ", Tarih: "2020-08-26", Koordinat: "39.400000,28.658200"},
	   		{EczaneAdi: "KÜBRA ECZANESİ", Telefon: "0236 515 22 82", Adres: "İHSANİYE MAH. 512 SOK. NO:35/A GÖLMARMARA", Ilce: "GÖLMARMARA", Tarih: "2020-08-26", Koordinat: "38.717594,27.911318"},
	   		{EczaneAdi: "ŞİFA ECZANESİ", Telefon: "0236 547 12 34", Adres: "CUMA MAH. ŞEHİT J.ASTS. B. ÜÇVÇ. ÖZGÜR ARTAR SOK. NO:2 GÖRDES", Ilce: "GÖRDES", Tarih: "2020-08-26", Koordinat: "38.932000,28.288800"},
	   		{EczaneAdi: "DERYA ECZANESİ", Telefon: "02365382732", Adres: "YENİ MAHALLE MENDERES CADDESİ 53 SOKAK NO:10 KIRKAĞAÇ", Ilce: "KIRKAĞAÇ", Tarih: "2020-08-26", Koordinat: "39.107900,27.675300"},
	   		{EczaneAdi: "AYDIN ECZANESİ", Telefon: "0236 571 28 65", Adres: "NAMIK KEMAL MH. CUMHURİYET CAD. NO:13 KÖPRÜBAŞI", Ilce: "KÖPRÜBAŞI", Tarih: "2020-08-26", Koordinat: "38.747700,28.403700"},
	   		{EczaneAdi: "DÖRT EYLÜL ECZANESİ", Telefon: "02368166343", Adres: "DEMİRCİLER CAD. NO:17 KULA", Ilce: "KULA", Tarih: "2020-08-26", Koordinat: "38.545300,28.645700"},
	   	}
	   	db.Create(&eczaneler)
	   	log.Println("Data inserted successfully.") */

	/* deleteQuery := `
			DELETE FROM pharmacies
	        WHERE id NOT IN (
	            SELECT id
	            FROM pharmacies
	            ORDER BY id
	            LIMIT 10
	        )`
		result := db.Exec(deleteQuery)
		if err := result.Error; err != nil {
			log.Fatal(err)
		}
		log.Println("Selected rows are deleted.") */
	return db
}
func main() {
	db := InitDB()
	router := mux.NewRouter()
	routers.RegisterRoutes(router, db)
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("error : %v", err)
	}
	log.Println("Server is running on port : ", router)
}
