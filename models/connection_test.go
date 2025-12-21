package models

import "testing"

func TestCreateAndUpdateConnectionInDb(t *testing.T) {
	ConnectTestingDatabase()
	MigrateDB(SqlDB)

	initial := Connection{
		HeaderID:        1,
		Timestamp:       "2024-01-01T00:00:00Z",
		Version:         "2.1.0",
		Manufacturer:    "Banana Republic",
		SerialNumber:    "1234",
		ConnectionState: "ONLINE",
	}

	err := CreateConnectionInDb(SqlDB, initial)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	var stored Connection
	if res := SqlDB.Where("serial_number = ?", initial.SerialNumber).First(&stored); res.Error != nil {
		t.Fatalf("expected connection in db: %v", res.Error)
	}

	updated := initial
	updated.ConnectionState = "OFFLINE"
	updated.LastNodeID = "node-1"
	updated.LastNodeSequenceID = 10
	updated.LastOrderID = "order-1"
	updated.LastOrderUpdateID = 2

	if err := UpdateConnectionInDb(SqlDB, stored, updated); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	var storedUpdated Connection
	if res := SqlDB.Where("serial_number = ?", initial.SerialNumber).First(&storedUpdated); res.Error != nil {
		t.Fatalf("expected updated connection in db: %v", res.Error)
	}

	if storedUpdated.ConnectionState != "OFFLINE" || storedUpdated.LastNodeID != "node-1" || storedUpdated.LastOrderID != "order-1" {
		t.Fatalf("unexpected updated values: %+v", storedUpdated)
	}
}
