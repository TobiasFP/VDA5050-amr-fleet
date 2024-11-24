package models

import "gorm.io/gorm"

type NodeState struct {
	gorm.Model
	NodeID          string "json:\"nodeId\""
	SequenceID      int    "json:\"sequenceId\""
	Released        bool   "json:\"released\""
	NodeDescription string "json:\"nodeDescription\""
}

type EdgeState struct {
	gorm.Model
	EdgeID          string `json:"edgeId"`
	SequenceID      int    `json:"sequenceId"`
	Released        bool   `json:"released"`
	EdgeDescription string `json:"edgeDescription"`
}

type ControlPoint struct {
	gorm.Model
	ControlPointID int
	X              float64 `json:"x"`
	Y              float64 `json:"y"`
	Weight         float64 `json:"weight"`
}

type ActionState struct {
	gorm.Model
	ActionID          string `json:"actionId"`
	ActionStatus      string `json:"actionStatus"`
	ActionType        string `json:"actionType"`
	ActionDescription string `json:"actionDescription"`
	ResultDescription string `json:"resultDescription"`
}

type BatteryState struct {
	gorm.Model
	BatteryCharge  float64 `json:"batteryCharge"`
	Charging       bool    `json:"charging"`
	BatteryVoltage float64 `json:"batteryVoltage"`
	BatteryHealth  float64 `json:"batteryHealth"`
	Reach          float64 `json:"reach"`
}

type StateError struct {
	gorm.Model
	ErrorType        string `json:"errorType"`
	ErrorLevel       string `json:"errorLevel"`
	ErrorDescription string `json:"errorDescription"`
	ErrorHint        string `json:"errorHint"`
}

type AgvPosition struct {
	gorm.Model
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"`
	PositionInitialized bool    `json:"positionInitialized"`
	MapDescription      string  `json:"mapDescription"`
	LocalizationScore   float64 `json:"localizationScore"`
	DeviationRange      float64 `json:"deviationRange"`
}

type Map struct {
	gorm.Model
	MapID          string `json:"mapId"`
	MapVersion     string `json:"mapVersion"`
	MapStatus      string `json:"mapStatus"`
	MapDescription string `json:"mapDescription"`
}

type SafetyState struct {
	gorm.Model
	EStop          string `json:"eStop"`
	FieldViolation bool   `json:"fieldViolation"`
}

type Velocity struct {
	gorm.Model
	Vx    float64 `json:"vx"`
	Vy    float64 `json:"vy"`
	Omega float64 `json:"omega"`
}

type Info struct {
	gorm.Model
	InfoType        string          `json:"infoType"`
	InfoLevel       string          `json:"infoLevel"`
	InfoReferences  []InfoReference `gorm:"many2many:info_info_references;" json:"infoReferences"`
	InfoDescription string          `json:"infoDescription"`
}

type InfoReference struct {
	gorm.Model
	ReferenceKey   string `json:"referenceKey"`
	ReferenceValue string `json:"referenceValue"`
}

type Load struct {
	gorm.Model
	LoadID       string `json:"loadId"`
	LoadType     string `json:"loadType"`
	LoadPosition string `json:"loadPosition"`
	// LoadDimensions LoadDimensions `json:"loadDimensions"`
	Weight float64 `json:"weight"`
}

// type LoadDimensions struct {
// 	gorm.Model
// 	Length float64 `json:"length"`
// 	Width  float64 `json:"width"`
// 	Height float64 `json:"height"`
// }

type State struct {
	gorm.Model
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
	BatteryState          BatteryState  `gorm:"column:battery_state_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"batteryState"`
	OperatingMode         string        `json:"operatingMode"`
	Errors                []StateError  `gorm:"many2many:state_errors;" json:"errors"`
	SafetyState           SafetyState   `gorm:"column:safety_state_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"safetyState"`
	Maps                  []Map         `gorm:"many2many:state_maps;" json:"maps"`
	ZoneSetID             string        `json:"zoneSetId"`
	Paused                bool          `json:"paused"`
	NewBaseRequest        bool          `json:"newBaseRequest"`
	DistanceSinceLastNode float64       `json:"distanceSinceLastNode"`
	AgvPosition           AgvPosition   `gorm:"column:agv_position_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agvPosition"`
	Velocity              Velocity      `gorm:"column:velocity_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"velocity"`
	Loads                 []Load        `gorm:"many2many:state_loads; "json:"loads"`
	Information           []Info        `gorm:"many2many:state_information;" json:"information"`
}
