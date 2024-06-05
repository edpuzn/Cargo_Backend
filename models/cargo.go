package models

import (
	"gorm.io/gorm"
)

type Cargo struct {
	gorm.Model
	SenderName      string  `json:"sender_name"`                    //türkçesi-> Gönderici Adı
	SenderAddress   string  `json:"sender_address"`                 //türkçesi-> Gönderici Adresi
	ReceiverName    string  `json:"receiver_name"`                  //türkçesi-> Alıcı Adı
	ReceiverAddress string  `json:"receiver_address"`               //türkçesi-> Alıcı Adresi
	Weight          float64 `json:"weight"`                         //türkçesi-> Ağırlık
	ShippingCost    float64 `json:"shipping_cost"`                  //türkçesi-> Kargo Ücreti
	TrackingNumber  string  `json:"tracking_number"  gorm:"unique"` //türkçesi-> Takip Numarası
	ShipmentStatus  string  `json:"shipment_status"`                //türkçesi-> Kargo Durumu
}
