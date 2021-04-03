package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/test_data"
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
