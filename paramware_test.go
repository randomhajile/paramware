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

	int64ParamName    string
	defaultInt64Value int64
	int64ParamHandler gin.HandlerFunc

	stringParamName    string
	defaultStringValue string
	stringParamHandler gin.HandlerFunc

	boolParamName    string
	defaultBoolValue bool
	boolParamHandler gin.HandlerFunc

	structParamName    string
	defaultStructValue structParam
	structParamHandler gin.HandlerFunc
}

type structParam struct {
	Value string
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

	suite.boolParamName = "boolParam"
	suite.defaultBoolValue = true
	suite.boolParamHandler = BoolParam(suite.boolParamName, suite.defaultBoolValue)

	suite.structParamName = "structParam"
	suite.defaultStructValue = structParam{Value: "default"}
	suite.structParamHandler = SetParam(
		suite.structParamName,
		suite.defaultStructValue,
		func(s string) (structParam, error) { return structParam{Value: s}, nil },
	)
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
	suite.EqualError(ctx.Errors[0], "expected int64 intParam param")
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

func (suite *paramsTestSuite) TestBoolParam() {
	testValue := !suite.defaultBoolValue
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=%t", suite.boolParamName, testValue),
		nil,
	)

	suite.boolParamHandler(ctx)
	suite.Equal(testValue, ctx.GetBool(suite.boolParamName))
}

func (suite *paramsTestSuite) TestBoolParamDefault() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	suite.boolParamHandler(ctx)
	suite.Equal(suite.defaultBoolValue, ctx.GetBool(suite.boolParamName))
}

func (suite *paramsTestSuite) TestBoolParamError() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=badvalue", suite.boolParamName),
		nil,
	)

	suite.boolParamHandler(ctx)
	suite.Len(ctx.Errors, 1)
	suite.EqualError(ctx.Errors[0], "expected bool boolParam param")
}

func (suite *paramsTestSuite) TestStruct() {
	testValue := "test-value"
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("?%s=%s", suite.structParamName, testValue),
		nil,
	)

	suite.structParamHandler(ctx)
	value := ctx.MustGet(suite.structParamName)
	suite.Equal(testValue, value.(structParam).Value)
}

func (suite *paramsTestSuite) TestStructParamDefault() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	suite.structParamHandler(ctx)
	value := ctx.MustGet(suite.structParamName)
	suite.Equal(suite.defaultStructValue, value)
}
