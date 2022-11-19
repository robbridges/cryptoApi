package controllers_test

import (
	"cryptoAPI/src/Router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetController(t *testing.T) {
	testRouter := Router.SetupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/crypto", nil)
	if err != nil {
		assert.Equal(t, 200, w.Code)
	}
	testRouter.ServeHTTP(w, req)

	assert.NotEqual(t, "", w.Body.String())
}
