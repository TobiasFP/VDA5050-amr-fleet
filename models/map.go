package models

import (
	"github.com/google/uuid"
)

type AmrMap struct {
	GormModelHiddenJson
	MapID          uuid.UUID  `gorm:"primarykey;unique;not null;type:uuid;" json:"mapId,omitempty"`
	MapVersion     string     `json:"mapVersion"`
	MapStatus      string     `json:"mapStatus"`
	MapDescription string     `json:"mapDescription"`
	MapDataId      int        `json:"-"`
	MapData        AmrMapData `gorm:"foreignKey:MapDataId;" json:"MapData"`
}

type AmrMapData struct {
	GormModelHiddenJson
	Data []byte `json:"Data"`
}
