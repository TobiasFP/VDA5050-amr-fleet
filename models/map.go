package models

import "gorm.io/gorm"

type AmrMap struct {
	gorm.Model
	MapID          string `gorm:"primarykey;unique;not null" json:"mapId"`
	MapVersion     string `json:"mapVersion"`
	MapStatus      string `json:"mapStatus"`
	MapDescription string `json:"mapDescription"`
}
