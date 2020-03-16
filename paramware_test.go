package paramware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type paramsTestSuite struct {
	suite.Suite

	int64ParamName     string
	defaultInt64Value  int64
	stringParamName    string
	defaultStringValue string

	int64ParamHandler  gin.HandlerFunc
	stringParamHandler gin.HandlerFunc
}

func TestParamsTestSuite(t *testing.T) {
	suite.Run(t, new(paramsTestSuite))
}

func (suite *paramsTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)

	suite.int64ParamName = "intParam"
	suite.defaultInt64Value = int64(10)
	suite.int64ParamHandler = Int64Param(suite.int64ParamName, suite.defaultInt64Value)

	suite.stringParamName = "stringParam"
	suite.defaultStringValue = "test"
	suite.stringParamHandler = StringParam(suite.stringParamName, suite.defaultStringValue)
}

func (suite *paramsTestSuite) TestInt64Param() {
	testValue := suite.defaultInt64Value + 1
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=%d", suite.int64ParamName, testValue),
		nil,
	)

	suite.int64ParamHandler(ctx)
	suite.Equal(testValue, ctx.GetInt64(suite.int64ParamName))
}

func (suite *paramsTestSuite) TestInt64ParamDefault() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	suite.int64ParamHandler(ctx)
	suite.Equal(suite.defaultInt64Value, ctx.GetInt64(suite.int64ParamName))
}

func (suite *paramsTestSuite) TestInt64ParamError() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=badvalue", suite.int64ParamName),
		nil,
	)

	suite.int64ParamHandler(ctx)
	suite.Len(ctx.Errors, 1)
	suite.EqualError(ctx.Errors[0], "expected int intParam param")
}

func (suite *paramsTestSuite) TestStringParam() {
	testValue := suite.defaultStringValue + "extra"
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=%s", suite.stringParamName, testValue),
		nil,
	)

	suite.stringParamHandler(ctx)
	suite.Equal(testValue, ctx.GetString(suite.stringParamName))
}

func (suite *paramsTestSuite) TestStringParamDefault() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	suite.stringParamHandler(ctx)
	suite.Equal(suite.defaultStringValue, ctx.GetString(suite.stringParamName))
}
