package controllers_test

import (
	controllers "cryptoAPI/src/controllers"
	_ "database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
}

func (s *UnitTestSuite) SetupTest() {
	viper.Set("URL", "localhost")
	viper.Set("PASSWORD", "postgres")
	viper.Set("USERNAME", "postgres")
}

func (s *UnitTestSuite) TestGetController() {
	s.SetupTest()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	controllers.GetCryptos(ctx)
	assert.Equal(s.T(), 200, w.Code)
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
