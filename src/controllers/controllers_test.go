package controllers_test

import (
	"bytes"
	router "cryptoAPI/src/Router"
	"cryptoAPI/src/models"
	_ "database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	r *gin.Engine
}

type Crypto models.Crypto

func (s *UnitTestSuite) SetupTest() {
	viper.Set("URL", "localhost")
	viper.Set("PASSWORD", "postgres")
	viper.Set("USERNAME", "postgres")
	s.r = router.SetupRouter()
}

func (s *UnitTestSuite) TestGetAllController() {
	s.SetupTest()

	req, err := http.NewRequest("GET", "/crypto/allcrypto", nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	var cryptos []Crypto
	err = json.Unmarshal(w.Body.Bytes(), &cryptos)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.NotEmpty(s.T(), cryptos)
}

func (s *UnitTestSuite) TestGetController() {
	s.SetupTest()

	req, err := http.NewRequest("GET", "/crypto/1", nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	testCryto := Crypto{1, "testcoin", 23.6, "www.fake.com"}
	var crypto Crypto
	err = json.Unmarshal(w.Body.Bytes(), &crypto)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), testCryto, crypto)
}

func (s *UnitTestSuite) TestGetController404() {
	s.SetupTest()
	expectedResponse := "Error: sql: no rows in result set"
	req, err := http.NewRequest("GET", "/crypto/478", nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	var response string
	err = json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
	assert.Equal(s.T(), expectedResponse, response)
}

func (s *UnitTestSuite) TestPostHander201() {
	s.SetupTest()

	cryptoExpected := Crypto{4, "fake2", 13.2, "www.fake2.com"}
	jsonValue, err := json.Marshal(cryptoExpected)
	req, err := http.NewRequest("POST", "/crypto/create?name=fake2&amount_owned=13.2&image_src=www.fake2.com", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	var expected Crypto
	err = json.Unmarshal(w.Body.Bytes(), &expected)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusCreated, w.Code)
	assert.Equal(s.T(), cryptoExpected.Name, expected.Name)
	assert.Equal(s.T(), cryptoExpected.Amount_Owned, expected.Amount_Owned)
	assert.Equal(s.T(), cryptoExpected.Image_Src, expected.Image_Src)

}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
