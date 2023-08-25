package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

type TestRequest struct {
	Method string
	Url    string
	Body   string
}

type TestResponse struct {
	StatusCode int
	Body       string
}

type TestCase struct {
	TestName    string
	Request     TestRequest
	HandlerFunc echo.HandlerFunc
	Response    TestResponse
}

func PrepareRequestsForHandlerTest(test *TestCase) (*http.Request, *httptest.ResponseRecorder) {
	// create a test request
	request := httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(""))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// create the recorder
	recorder := httptest.NewRecorder()

	return request, recorder
}
