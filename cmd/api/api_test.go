package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateRide(t *testing.T) {
	t.Run("should return a ride price", func(t *testing.T) {
		app := Build()
		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-01T10:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, 200)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		assert.Equal(t, `{"price":21}`, string(resBody))
	})

	t.Run("should return a 422 response if date is not valid", func(t *testing.T) {
		app := Build()
		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-0110:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, 422)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		assert.Equal(t, `{"message":"Invalid Date"}`, string(resBody))
	})
}
