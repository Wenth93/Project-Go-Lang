package handlers

import (
	"net/http"
	"testing"

	"github.com/Wenth93/Project-Go-Lang/tests"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	cases := []tests.TestCase{
		{
			TestName: "test /live route",
			Request: tests.TestRequest{
				Method: http.MethodGet,
				Url:    "/live",
			},
			HandlerFunc: func(c echo.Context) error {
				return NewHealthHandler().IsAlive(c)
			},
			Response: tests.TestResponse{
				StatusCode: 200,
				Body: `{
					"message": "Ok"
				}`,
			},
		},
	}

	s := echo.New()
	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request, recorder := tests.PrepareRequestsForHandlerTest(&test)

			// create new context
			newContext := s.NewContext(request, recorder)

			// handler function will be called on this new context (mocked context)
			if assert.NoError(t, test.HandlerFunc(newContext)) {
				assert.Equal(t, test.Response.StatusCode, recorder.Code)
				assert.JSONEq(t, test.Response.Body, recorder.Body.String())
			}
		})
	}
}
