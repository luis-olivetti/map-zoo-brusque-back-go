package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type mockUserService struct {
	mock.Mock
}

func (s *mockUserService) Authenticate(username, password string) (bool, error) {
	args := s.Called(username, password)
	return args.Bool(0), args.Error(1)
}

func (s *mockUserService) GenerateJWT(username string) (*model.UserJWT, error) {
	args := s.Called(username)
	return args.Get(0).(*model.UserJWT), args.Error(1)
}

func TestUserHandler_Login(t *testing.T) {
	// Arrange
	recorder := httptest.NewRecorder()
	ctx := getTestGinContext(recorder)

	userServiceMock := new(mockUserService)
	userServiceMock.On("Authenticate", "test", "test").Return(true, nil)

	fakeJWT := &model.UserJWT{
		Token: "fakeJWT",
	}
	userServiceMock.On("GenerateJWT", "test").Return(fakeJWT, nil)

	handler := NewUserHandler(nil, userServiceMock)

	jsonBody := `{"username": "test", "password": "test"}`

	makePost(ctx, jsonBody)

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedJSON := `{"message":"success","data":{"token":"fakeJWT"}}`
	assert.Equal(t, expectedJSON, recorder.Body.String())
}

func TestUserHandler_Login_Unauthorized(t *testing.T) {
	// Arrange
	recorder := httptest.NewRecorder()
	ctx := getTestGinContext(recorder)

	handlerMock := &Handler{
		logger: createNoOpLogger(),
	}

	userServiceMock := new(mockUserService)
	userServiceMock.On("Authenticate", "test", "test").Return(false, nil)

	fakeJWT := &model.UserJWT{
		Token: "fakeJWT",
	}
	userServiceMock.On("GenerateJWT", "test").Return(fakeJWT, nil)

	handler := NewUserHandler(handlerMock, userServiceMock)

	jsonBody := `{"username": "test", "password": "test"}`
	makePost(ctx, jsonBody)

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)

	expectedJSON := `{"message":"Unauthorized","data":{}}`
	assert.Equal(t, expectedJSON, recorder.Body.String())
}

func TestUserHandler_Login_InvalidRequest(t *testing.T) {
	// Arrange
	recorder := httptest.NewRecorder()
	ctx := getTestGinContext(recorder)

	handlerMock := &Handler{
		logger: createNoOpLogger(),
	}

	userServiceMock := new(mockUserService)
	userServiceMock.On("Authenticate", "test", "test").Return(false, nil)

	fakeJWT := &model.UserJWT{
		Token: "fakeJWT",
	}
	userServiceMock.On("GenerateJWT", "test").Return(fakeJWT, nil)

	handler := NewUserHandler(handlerMock, userServiceMock)

	invalidJsonBody := "invalid JSON"
	makePost(ctx, invalidJsonBody)

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	expectedJSON := `{"message":"Bad Request","data":{}}`
	assert.Equal(t, expectedJSON, recorder.Body.String())
}

func getTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func makePost(ctx *gin.Context, jsonBody string) {
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Body = io.NopCloser(strings.NewReader(jsonBody))
}

func createNoOpLogger() *log.Logger {
	logger := zap.NewNop()
	return &log.Logger{Logger: logger}
}
