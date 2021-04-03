package apis

import (
	"bytes"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/test_data"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiTestCase struct {
	tag              string
	method           string
	urlToServe       string
	urlToHit         string
	body             string
	headers          []requestHeader
	function         gin.HandlerFunc
	status           int
	responseFilePath string
}

type requestHeader struct {
	key   string
	value string
}

// Creates new router in testing mode
func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	config.Config.DB = test_data.ResetDB()

	return router
}

// Used to run single API test case. It makes HTTP request and returns its response
func testAPI(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, body string, headers []requestHeader) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	for _, header := range headers {
		req.Header.Add(header.key, header.value)
	}
	router.ServeHTTP(res, req)
	return res
}

// Used to run suite (list) of test cases. It checks JSON response is same as expected data in test case file.
// All test expected test case responses are stored in `test_data/test_case_data` folder in format `<suite_name>_t<number>.json`
func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.urlToServe, test.urlToHit, test.function, test.body, test.headers)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.responseFilePath != "" {
			response, _ := ioutil.ReadFile(test.responseFilePath)
			assert.JSONEq(t, string(response), res.Body.String(), test.tag)
		}
	}
}
