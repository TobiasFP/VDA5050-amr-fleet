package models

// As this type of data is much better handled in a nosql context, we will use elasticfor this data.
type ActionParameter struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// These will live in the SQL database
type Action struct {
	GormModelHiddenJson

	ActionID          string            `json:"actionId"`
	ActionType        string            `json:"actionType"`
	BlockingType      string            `json:"blockingType"`
	ActionDescription string            `json:"actionDescription"`
	ActionParameters  []ActionParameter `gorm:"-" json:"actionParameters"`
}

type InstantAction struct {
	GormModelHiddenJson
	HeaderID     int      `json:"headerId"`
	Timestamp    string   `json:"timestamp"`
	Version      string   `json:"version"`
	Manufacturer string   `json:"manufacturer"`
	SerialNumber string   `json:"serialNumber"`
	Actions      []Action `gorm:"many2many:instantaction_actions;" json:"actions"`
}
