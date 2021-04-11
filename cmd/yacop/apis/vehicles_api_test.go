package apis

import (
	"bytes"
	"encoding/json"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/fitchlol/yacop/cmd/yacop/test_data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestVehicles(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/vehicles"
	var emptyHeaders []requestHeader
	emptyBody := ""

	runAPITests(t, []apiTestCase{
		{
			"get /vehicles",
			"GET",
			"/vehicles",
			"/vehicles",
			emptyBody,
			emptyHeaders,
			Vehicles,
			http.StatusOK,
			path + "/all.json",
		},
	})
}

func TestVehicleCreate(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/vehicles"

	vehicleCreateContent, err := ioutil.ReadFile(path + "/requests/vehicles_create_new_vehicle.json")
	if err != nil {
		panic(err)
	}
	vehicleCreateData := new(bytes.Buffer)

	if err := json.Compact(vehicleCreateData, vehicleCreateContent); err != nil {
		panic(err)
	}
	vehicleCreateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(router, "POST", "/vehicles", "/vehicles", VehicleCreate, vehicleCreateData.String(), vehicleCreateHeaders)
	assert.Equal(t, http.StatusCreated, res.Code, "create new vehicle")
	var createdVehicle models.Vehicle
	err = json.Unmarshal([]byte(res.Body.String()), &createdVehicle)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, createdVehicle.WeightInKilograms)
	assert.NotNil(t, createdVehicle.MaximumKilowatts)
	assert.NotNil(t, createdVehicle.MaximumKilometersPerHour)
	assert.NotNil(t, createdVehicle.FullName)
	assert.NotNil(t, createdVehicle.Manufacturer)
}

func TestVehicleUpdate(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/vehicles"

	vehicleUpdateContent, err := ioutil.ReadFile(path + "/requests/vehicles_update.json")
	if err != nil {
		panic(err)
	}
	vehicleUpdateData := new(bytes.Buffer)

	if err := json.Compact(vehicleUpdateData, vehicleUpdateContent); err != nil {
		panic(err)
	}
	vehicleUpdateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(
		router,
		"PUT",
		"/vehicles/:id",
		"/vehicles/00000000-0000-0000-0000-000000000001",
		VehicleUpdate,
		vehicleUpdateData.String(),
		vehicleUpdateHeaders,
	)
	assert.Equal(t, http.StatusOK, res.Code, "update existing vehicle")
	var updatedVehicle models.Vehicle
	err = json.Unmarshal([]byte(res.Body.String()), &updatedVehicle)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "A6 3.0 TDI", updatedVehicle.FullName)
}

func TestVehicleUpdateUnknown(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/vehicles"

	vehicleUpdateContent, err := ioutil.ReadFile(path + "/requests/vehicles_update.json")
	if err != nil {
		panic(err)
	}
	vehicleUpdateData := new(bytes.Buffer)

	if err := json.Compact(vehicleUpdateData, vehicleUpdateContent); err != nil {
		panic(err)
	}
	vehicleUpdateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(
		router,
		"PUT",
		"/vehicles/:id",
		"/vehicles/10000000-0000-0000-0000-000000000000",
		VehicleUpdate,
		vehicleUpdateData.String(),
		vehicleUpdateHeaders,
	)
	assert.Equal(t, http.StatusNotFound, res.Code, "update unknown vehicle")
}

func TestVehicleDelete(t *testing.T) {
	router := newRouter()
	var emptyHeaders []requestHeader
	res := testAPI(
		router,
		"DELETE",
		"/vehicles/:id",
		"/vehicles/00000000-0000-0000-0000-000000000001",
		VehicleDelete,
		"",
		emptyHeaders,
	)
	assert.Equal(t, http.StatusNoContent, res.Code, "delete existing vehicle")
	assert.Equal(t, "", res.Body.String())
}

func TestVehicleDeleteUnknown(t *testing.T) {
	router := newRouter()
	var emptyHeaders []requestHeader
	res := testAPI(
		router,
		"DELETE",
		"/vehicles/:id",
		"/vehicles/10000000-0000-0000-0000-000000000000",
		VehicleDelete,
		"",
		emptyHeaders,
	)
	assert.Equal(t, http.StatusNotFound, res.Code, "delete unknown vehicle")
}
