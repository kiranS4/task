package parse

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite
	data PareseIn
}

func (suite *MySuite) SetupTest() {
	suite.data = NewData("*/15 0 1,15 * 1-5 /usr/bin/find")
	suite.data.Parse()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}

func (suite *MySuite) TestHour() {
	suite.Equal("0", suite.data.GetHour())
}

func (suite *MySuite) TestDayofMonth() {
	suite.Equal("1 15", suite.data.GetDayOfMonth())
}

func (suite *MySuite) TestMin() {
	suite.Equal("0 15 30 45 ", suite.data.GetMin())
}
