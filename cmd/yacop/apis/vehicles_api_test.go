package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/test_data"
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
