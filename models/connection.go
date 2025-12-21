package models

import "gorm.io/gorm"

type Connection struct {
	GormModelHiddenJson
	HeaderID           int    `json:"headerId"`
	Timestamp          string `json:"timestamp"`
	Version            string `json:"version"`
	Manufacturer       string `json:"manufacturer"`
	SerialNumber       string `json:"serialNumber"`
	ConnectionState    string `json:"connectionState"`
	LastNodeID         string `json:"lastNodeId"`
	LastNodeSequenceID int    `json:"lastNodeSequenceId"`
	LastOrderID        string `json:"lastOrderId"`
	LastOrderUpdateID  int    `json:"lastOrderUpdateId"`
}

func CreateConnectionInDb(db *gorm.DB, conn Connection) error {
	res := db.Create(&conn)
	return res.Error
}

func UpdateConnectionInDb(db *gorm.DB, connInDb Connection, connFromMqtt Connection) error {
	connFromMqtt.ID = connInDb.ID
	res := db.Save(&connFromMqtt)
	return res.Error
}
