package models

import (
	"gorm.io/gorm"
)

type NodeState struct {
	GormModelHiddenJson
	NodeID          string `json:"nodeId"`
	SequenceID      int    `json:"sequenceId"`
	Released        bool   `json:"released"`
	NodeDescription string `json:"nodeDescription"`
}

type EdgeState struct {
	GormModelHiddenJson
	EdgeID          string `json:"edgeId"`
	SequenceID      int    `json:"sequenceId"`
	Released        bool   `json:"released"`
	EdgeDescription string `json:"edgeDescription"`
}

type ActionState struct {
	GormModelHiddenJson
	ActionID          string `json:"actionId"`
	ActionStatus      string `json:"actionStatus"`
	ActionType        string `json:"actionType"`
	ActionDescription string `json:"actionDescription"`
	ResultDescription string `json:"resultDescription"`
}

type BatteryState struct {
	GormModelHiddenJson
	BatteryCharge  float64 `json:"batteryCharge"`
	Charging       bool    `json:"charging"`
	BatteryVoltage float64 `json:"batteryVoltage"`
	BatteryHealth  float64 `json:"batteryHealth"`
	Reach          float64 `json:"reach"`
}

type StateError struct {
	GormModelHiddenJson
	ErrorType        string `json:"errorType"`
	ErrorLevel       string `json:"errorLevel"`
	ErrorDescription string `json:"errorDescription"`
	ErrorHint        string `json:"errorHint"`
}

type AgvPosition struct {
	GormModelHiddenJson
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"` // We should, with gorm, relate this id to the actual map
	PositionInitialized bool    `json:"positionInitialized"`
	MapDescription      string  `json:"mapDescription"`
	LocalizationScore   float64 `json:"localizationScore"`
	DeviationRange      float64 `json:"deviationRange"`
}

type SafetyState struct {
	GormModelHiddenJson
	EStop          string `json:"eStop"`
	FieldViolation bool   `json:"fieldViolation"`
}

type Velocity struct {
	GormModelHiddenJson
	Vx    float64 `json:"vx"`
	Vy    float64 `json:"vy"`
	Omega float64 `json:"omega"`
}

type Info struct {
	GormModelHiddenJson
	InfoType        string          `json:"infoType"`
	InfoLevel       string          `json:"infoLevel"`
	InfoReferences  []InfoReference `gorm:"many2many:info_info_references;" json:"infoReferences"`
	InfoDescription string          `json:"infoDescription"`
}

type InfoReference struct {
	GormModelHiddenJson
	ReferenceKey   string `json:"referenceKey"`
	ReferenceValue string `json:"referenceValue"`
}

type Load struct {
	GormModelHiddenJson
	LoadID       string `json:"loadId"`
	LoadType     string `json:"loadType"`
	LoadPosition string `json:"loadPosition"`
	// LoadDimensions LoadDimensions `json:"loadDimensions"`
	Weight float64 `json:"weight"`
}

// type LoadDimensions struct {
// 	GormModelHiddenJson
// 	Length float64 `json:"length"`
// 	Width  float64 `json:"width"`
// 	Height float64 `json:"height"`
// }

type State struct {
	GormModelHiddenJson
	HeaderID              int           `json:"headerId"`
	Timestamp             string        `json:"timestamp"`
	Version               string        `json:"version"`
	Manufacturer          string        `json:"manufacturer"`
	SerialNumber          string        `json:"serialNumber"`
	OrderID               string        `json:"orderId"`
	OrderUpdateID         int           `json:"orderUpdateId"`
	LastNodeID            string        `json:"lastNodeId"`
	LastNodeSequenceID    int           `json:"lastNodeSequenceId"`
	NodeStates            []NodeState   `gorm:"many2many:state_node_state;" json:"nodeStates"`
	EdgeStates            []EdgeState   `gorm:"many2many:state_edge_state;" json:"edgeStates"`
	Driving               bool          `json:"driving"`
	ActionStates          []ActionState `gorm:"many2many:state_action_state;" json:"actionStates"`
	BatteryStateID        int           `json:"-"` // Not  in the vda55 struct, simply a field for GORM
	BatteryState          BatteryState  `gorm:"foreignKey:BatteryStateID;" json:"batteryState"`
	OperatingMode         string        `json:"operatingMode"`
	Errors                []StateError  `gorm:"many2many:state_errors;" json:"errors"`
	SafetyStateID         int           `json:"-"` // Not  in the vda55 struct, simply a field for GORM
	SafetyState           SafetyState   `gorm:"foreignKey:SafetyStateID;" json:"safetyState"`
	Maps                  []AmrMap      `gorm:"many2many:state_maps;" json:"maps"`
	ZoneSetID             string        `json:"zoneSetId"`
	Paused                bool          `json:"paused"`
	NewBaseRequest        bool          `json:"newBaseRequest"`
	DistanceSinceLastNode float64       `json:"distanceSinceLastNode"`
	AgvPositionID         int           `json:"-"` // Not  in the vda55 struct, simply a field for GORM
	AgvPosition           AgvPosition   `gorm:"foreignKey:AgvPositionID;" json:"agvPosition"`
	Velocity              Velocity      `gorm:"-; " json:"velocity"` // We should not save such volatile data in our database. This should only be extracted from mqtt.
	Loads                 []Load        `gorm:"many2many:state_loads;" json:"loads"`
	Information           []Info        `gorm:"many2many:state_information;" json:"information"`
}

func CreateAmrStateInDb(db *gorm.DB, state State) error {
	// Since we know that the amr is not in the db,
	// we can create empty battery and map states
	batteryCreateRes := db.Create(&state.BatteryState)

	if batteryCreateRes.Error != nil {
		return batteryCreateRes.Error
	}

	safetyStateCreateRes := db.Create(&state.SafetyState)

	if safetyStateCreateRes.Error != nil {
		return safetyStateCreateRes.Error
	}

	res := db.Create(&state)
	return res.Error

}

func UpdateAmrStateInDb(db *gorm.DB, amrInDB State, amrFromMqtt State) error {
	amrFromMqtt.AgvPosition.ID = uint(amrInDB.AgvPositionID)
	AgvPositionRes := db.Save(&amrFromMqtt.AgvPosition)
	if AgvPositionRes.Error != nil {
		return AgvPositionRes.Error
	}
	amrFromMqtt.BatteryState.ID = uint(amrInDB.BatteryStateID)
	BatteryStateRes := db.Save(&amrFromMqtt.BatteryState)
	if BatteryStateRes.Error != nil {
		return BatteryStateRes.Error
	}

	amrFromMqtt.SafetyState.ID = uint(amrInDB.SafetyStateID)
	SafetyStateRes := db.Save(&amrFromMqtt.SafetyState)
	if SafetyStateRes.Error != nil {
		return SafetyStateRes.Error
	}
	amrFromMqtt.ID = amrInDB.ID
	amrSaveRes := db.Save(&amrFromMqtt)
	return amrSaveRes.Error
}
