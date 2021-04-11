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

func TestManufacturers(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/manufacturers"
	var emptyHeaders []requestHeader
	emptyBody := ""

	runAPITests(t, []apiTestCase{
		{
			"get /manufacturers",
			"GET",
			"/manufacturers",
			"/manufacturers",
			emptyBody,
			emptyHeaders,
			Manufacturers,
			http.StatusOK,
			path + "/all.json",
		},
	})
}

func TestManufacturerCreate(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/manufacturers"

	manufacturerCreateContent, err := ioutil.ReadFile(path + "/requests/manufacturers_create_new_manufacturer.json")
	if err != nil {
		panic(err)
	}
	manufacturerCreateData := new(bytes.Buffer)

	if err := json.Compact(manufacturerCreateData, manufacturerCreateContent); err != nil {
		panic(err)
	}
	manufacturerCreateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(router, "POST", "/manufacturers", "/manufacturers", ManufacturerCreate, manufacturerCreateData.String(), manufacturerCreateHeaders)
	assert.Equal(t, http.StatusCreated, res.Code, "create new manufacturer")
	var createdManufacturer models.Manufacturer
	err = json.Unmarshal([]byte(res.Body.String()), &createdManufacturer)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, createdManufacturer.Name)
}

func TestManufacturerUpdate(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/manufacturers"

	manufacturerUpdateContent, err := ioutil.ReadFile(path + "/requests/manufacturers_update.json")
	if err != nil {
		panic(err)
	}
	manufacturerUpdateData := new(bytes.Buffer)

	if err := json.Compact(manufacturerUpdateData, manufacturerUpdateContent); err != nil {
		panic(err)
	}
	manufacturerUpdateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(
		router,
		"PUT",
		"/manufacturers/:id",
		"/manufacturers/00000000-0000-0000-0000-000000000010",
		ManufacturerUpdate,
		manufacturerUpdateData.String(),
		manufacturerUpdateHeaders,
	)
	assert.Equal(t, http.StatusOK, res.Code, "update existing manufacturer")
	var updatedManufacturer models.Manufacturer
	err = json.Unmarshal([]byte(res.Body.String()), &updatedManufacturer)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "BMW", updatedManufacturer.Name)
}

func TestManufacturerUpdateUnknown(t *testing.T) {
	path := test_data.GetTestCaseFolder() + "/manufacturers"

	manufacturerUpdateContent, err := ioutil.ReadFile(path + "/requests/manufacturers_update.json")
	if err != nil {
		panic(err)
	}
	manufacturerUpdateData := new(bytes.Buffer)

	if err := json.Compact(manufacturerUpdateData, manufacturerUpdateContent); err != nil {
		panic(err)
	}
	manufacturerUpdateHeaders := []requestHeader{{key: "Content-Type", value: "application/json"}}

	router := newRouter()
	res := testAPI(
		router,
		"PUT",
		"/manufacturers/:id",
		"/manufacturers/10000000-0000-0000-0000-000000000000",
		ManufacturerUpdate,
		manufacturerUpdateData.String(),
		manufacturerUpdateHeaders,
	)
	assert.Equal(t, http.StatusNotFound, res.Code, "update unknown manufacturer")
}

func TestManufacturerDelete(t *testing.T) {
	router := newRouter()
	var emptyHeaders []requestHeader
	res := testAPI(
		router,
		"DELETE",
		"/manufacturers/:id",
		"/manufacturers/00000000-0000-0000-0000-000000000010",
		ManufacturerDelete,
		"",
		emptyHeaders,
	)
	assert.Equal(t, http.StatusNoContent, res.Code, "delete existing manufacturer")
	assert.Equal(t, "", res.Body.String())
}

func TestManufacturerDeleteUnknown(t *testing.T) {
	router := newRouter()
	var emptyHeaders []requestHeader
	res := testAPI(
		router,
		"DELETE",
		"/manufacturers/:id",
		"/manufacturers/10000000-0000-0000-0000-000000000000",
		ManufacturerDelete,
		"",
		emptyHeaders,
	)
	assert.Equal(t, http.StatusNotFound, res.Code, "delete unknown manufacturer")
}
