package models

import (
	"testing"
)

func TestUpdateAmrStateInDb(t *testing.T) {
	ConnectTestingDatabase()
	MigrateDB(SqlDB)
	state := State{
		GormModelHiddenJson: GormModelHiddenJson{},
		HeaderID:            0,
		Timestamp:           "",
		Version:             "",
		Manufacturer:        "",
		SerialNumber:        "",
		OrderID:             "",
		OrderUpdateID:       0,
		LastNodeID:          "",
		LastNodeSequenceID:  0,
		NodeStates:          []NodeState{},
		EdgeStates:          []EdgeState{},
		Driving:             false,
		ActionStates:        []ActionState{},
		BatteryStateID:      0,
		BatteryState: BatteryState{
			GormModelHiddenJson: GormModelHiddenJson{},
			BatteryCharge:       87,
			Charging:            true,
			BatteryVoltage:      4,
			BatteryHealth:       100,
			Reach:               0,
		},
		OperatingMode: "",
		Errors:        []StateError{},
		SafetyStateID: 0,
		SafetyState: SafetyState{
			GormModelHiddenJson: GormModelHiddenJson{},
			EStop:               "false",
			FieldViolation:      false,
		},
		Maps:                  []AmrMap{},
		ZoneSetID:             "",
		Paused:                false,
		NewBaseRequest:        false,
		DistanceSinceLastNode: 0,
		AgvPositionID:         0,
		AgvPosition:           AgvPosition{},
		Velocity:              Velocity{},
		Loads:                 []Load{},
		Information:           []Info{},
	}
	err := CreateAmrStateInDb(SqlDB, state)

	if err != nil {
		t.Errorf("got %q", err.Error())
	}
}
