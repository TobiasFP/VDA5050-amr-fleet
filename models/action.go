package models

type ActionParameter struct {
	GormModelHiddenJson

	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type Action struct {
	GormModelHiddenJson

	ActionID          string            `json:"actionId"`
	ActionType        string            `json:"actionType"`
	BlockingType      string            `json:"blockingType"`
	ActionDescription string            `json:"actionDescription"`
	ActionParameters  []ActionParameter `gorm:"many2many:action_actionparameters;" json:"actionParameters"`
}

type InstantAction struct {
	GormModelHiddenJson
	HeaderID     int      `json:"headerId"`
	Timestamp    string   `json:"timestamp"`
	Version      string   `json:"version"`
	Manufacturer string   `json:"manufacturer"`
	SerialNumber string   `json:"serialNumber"`
	Actions      []Action `json:"actions"`
}
