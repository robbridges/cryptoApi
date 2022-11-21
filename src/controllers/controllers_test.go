package controllers_test

import (
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

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
