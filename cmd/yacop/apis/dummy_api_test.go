package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/test_data"
	"net/http"
	"testing"
)

func TestUsers(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	var emptyHeaders []requestHeader
	emptyBody := ""

	runAPITests(t, []apiTestCase{
		{
			"get /dummy/hello-world",
			"GET",
			"/dummy/hello-world",
			"/dummy/hello-world",
			emptyBody,
			emptyHeaders,
			SayHelloWorld,
			http.StatusOK,
			path + "/dummy_hello-world.json",
		},
	})
}
