package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type mockUserService struct{}

func (s mockUserService) Authenticate(username, password string) (bool, error) {
	return true, nil
}

func (s mockUserService) GenerateJWT(username string) (*model.UserJWT, error) {
	return &model.UserJWT{
		Token: "",
	}, nil
}

func TestUserHandler_Login(t *testing.T) {
	// Arrange
	recorder := httptest.NewRecorder()
	ctx := getTestGinContext(recorder)

	handler := NewUserHandler(nil, mockUserService{})

	jsonBody := `{"username": "test", "password": "test"}`

	makePost(ctx, jsonBody)

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusOK, recorder.Code)
	// assert.Contains(t, recorder.Body.String(), "expected_response_content")
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
