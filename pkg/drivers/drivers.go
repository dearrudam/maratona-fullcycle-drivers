package drivers

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

func init() {
	_ = godotenv.Load()
}

// Driver struct represents a driver
type Driver struct {
	ID   string `json:"uuid"`
	Name string `json:"name"`
}

// MarshallToJSON method marshalls such one in a []byte
func (driver Driver) MarshallToJSON() []byte {
	rawData, _ := json.Marshal(driver)
	return rawData
}

// Drivers represents a source of available drivers
type Drivers struct {
	DriversList []Driver `json:"drivers"`
}

// GetDriverByID if a method for retrieve a Driver for a given ID
func (source Drivers) GetDriverByID(id string) (Driver, error) {

	for _, driver := range source.DriversList {
		if driver.ID == id {
			return driver, nil
		}
	}

	return Driver{}, fmt.Errorf("driver %v not found", id)
}

// UnmarshalFromJSON function retrieves a Drivers reference based on a []byte readed from a json file
func UnmarshalFromJSON(rawData []byte) (Drivers, error) {
	var drivers Drivers
	if err := json.Unmarshal(rawData, &drivers); err != nil {
		return drivers, err
	}
	return drivers, nil
}

// MarshallToJSON method marshalls such one in a []byte
func (source Drivers) MarshallToJSON() []byte {
	rawData, _ := json.Marshal(source)
	return rawData
}

// LoadDrivers returns a Drivers loaded from 'drivers.json'
func LoadDrivers() Drivers {
	sourceFile := "drivers.json"
	if os.Getenv("DRIVERS_SOURCE") != "" {
		sourceFile = os.Getenv("DRIVERS_SOURCE")
	}
	jsonFile, err := os.Open(sourceFile)
	if err != nil {
		panic(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err.Error())
	}

	drivers, err := UnmarshalFromJSON(data)
	if err != nil {
		panic(err.Error())
	}

	return drivers
}
