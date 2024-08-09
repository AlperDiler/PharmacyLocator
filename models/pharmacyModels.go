package models

type Pharmacy struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	EczaneAdi string  `gorm:"eczaneAdi" json:"eczaneAdi"`
	Telefon   string  `gorm:"telefon" json:"telefon"`
	Adres     string  `gorm:"adres" json:"adres"`
	Ilce      string  `gorm:"ilce" json:"ilce"`
	Tarih     string  `gorm:"tarih" json:"tarih"`
	Koordinat string  `gorm:"koorinat" json:"koorinat"`
	Latidute  float64 `gorm:"latidute" json:"latidute"`
	Longidute float64 `gorm:"longidute" json:"longidute"`
}

type UserCoords struct {
	UserLat string `gorm:"userLat" json:"userLat"`
	UserLon string `gorm:"userLon" json:"userLon"`
}
