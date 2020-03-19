package drivers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalFromJSON(t *testing.T) {
	rawData := []byte(`
	{
		"drivers": [
			{
				"uuid": "45688cd6-7a27-4a7b-89c5-a9b604eefe2f",
				"name": "Wesley"
			},
			{
				"uuid": "9a118e4d-821a-44c7-accc-fa99ac4be01a",
				"name": "Luiz"
			}
		]
	}
	`)

	expectedDrivers := Drivers{
		DriversList: []Driver{
			Driver{ID: "45688cd6-7a27-4a7b-89c5-a9b604eefe2f", Name: "Wesley"},
			Driver{ID: "9a118e4d-821a-44c7-accc-fa99ac4be01a", Name: "Luiz"},
		},
	}

	drivers, err := UnmarshalFromJSON(rawData)
	if err != nil {
		assert.Fail(t, "failure to unmarshall the json: %v", err)
	}

	assert.Equal(t, len(expectedDrivers.DriversList), len(drivers.DriversList), "size of loaded drivers is wrong")

	for index, expectedDriver := range expectedDrivers.DriversList {
		assert.Equal(t, expectedDriver.ID, drivers.DriversList[index].ID, "driver's ID from the loaded drivers [%v] not expected", index)
		assert.Equal(t, expectedDriver.Name, drivers.DriversList[index].Name, "driver's ID from the loaded drivers [%v] not expected", index)
	}

}

func TestMarshalToJSON(t *testing.T) {

	expectedRawData := []byte(`{"drivers":[{"uuid":"45688cd6-7a27-4a7b-89c5-a9b604eefe2f","name":"Wesley"}]}`)

	actualDrivers := Drivers{
		DriversList: []Driver{
			Driver{ID: "45688cd6-7a27-4a7b-89c5-a9b604eefe2f", Name: "Wesley"},
		},
	}

	actualRawData := actualDrivers.MarshallToJSON()

	assert.Equal(t, expectedRawData, actualRawData, "Drivers's marshall to json is not working properly")

	expectedRawData = []byte(`{"uuid":"45688cd6-7a27-4a7b-89c5-a9b604eefe2f","name":"Wesley"}`)
	actualRawData = actualDrivers.DriversList[0].MarshallToJSON()

	assert.Equal(t, expectedRawData, actualRawData, "Driver's marshall to json is not working properly")

}

func TestLoadDrivers(t *testing.T) {
	expectedDrivers := Drivers{
		DriversList: []Driver{
			Driver{ID: "45688cd6-7a27-4a7b-89c5-a9b604eefe2f", Name: "Wesley"},
			Driver{ID: "9a118e4d-821a-44c7-accc-fa99ac4be01a", Name: "Luiz"},
		},
	}

	os.Setenv("DRIVERS_SOURCE","./../../drivers.json")

	drivers := LoadDrivers()

	assert.Equal(t, len(expectedDrivers.DriversList), len(drivers.DriversList), "size of loaded drivers is wrong")

	for index, expectedDriver := range expectedDrivers.DriversList {
		assert.Equal(t, expectedDriver.ID, drivers.DriversList[index].ID, "driver's ID from the loaded drivers [%v] not expected", index)
		assert.Equal(t, expectedDriver.Name, drivers.DriversList[index].Name, "driver's ID from the loaded drivers [%v] not expected", index)
	}
}

func TestGetDriverByID(t *testing.T) {
	drivers := Drivers{
		DriversList: []Driver{
			Driver{ID: "45688cd6-7a27-4a7b-89c5-a9b604eefe2f", Name: "Wesley"},
			Driver{ID: "9a118e4d-821a-44c7-accc-fa99ac4be01a", Name: "Luiz"},
		},
	}

	id := "45688cd6-7a27-4a7b-89c5-a9b604eefe2f"
	foundDriver, err := drivers.GetDriverByID(id)
	if err != nil {
		assert.Fail(t, "cannot found the existent driver by id %v", id)
	}
	expectedDriver := drivers.DriversList[0]
	assert.Equal(t, expectedDriver.ID, foundDriver.ID, "retrieved an unexpected driver by id %v", id)
	assert.Equal(t, expectedDriver.Name, foundDriver.Name, "retrieved an unexpected driver by id %v", id)

	id = "9a118e4d-821a-44c7-accc-fa99ac4be01a"
	foundDriver, err = drivers.GetDriverByID(id)
	if err != nil {
		assert.Fail(t, "cannot found the existent driver by id %v", id)
	}
	expectedDriver = drivers.DriversList[1]
	assert.Equal(t, expectedDriver.ID, foundDriver.ID, "retrieved an unexpected driver by id %v", id)
	assert.Equal(t, expectedDriver.Name, foundDriver.Name, "retrieved an unexpected driver by id %v", id)

	foundDriver, err = drivers.GetDriverByID("erwerwerwerw")
	if err == nil {
		assert.Fail(t, "must return error for a invalid id")
	}
}
