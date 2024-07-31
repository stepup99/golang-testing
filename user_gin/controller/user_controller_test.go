package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stepup99/golang-testing/user_gin/domain"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	GetUserFunc func(userId int) (*domain.User, error)
}

func (m *MockUserService) GetUser(userId int) (*domain.User, error) {
	return m.GetUserFunc(userId)
}

func TestGetUser(t *testing.T) {
	t.Run("getUser success", func(t *testing.T) {
		UserService = &MockUserService{
			GetUserFunc: func(userId int) (*domain.User, error) {
				fmt.Println("27 ????>>>>>>>>>>>>>>>  ")
				return &domain.User{ID: userId, Name: "John Doe"}, nil
			},
		}

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "/user/1", nil)
		c.Params = gin.Params{
			{Key: "user_id", Value: "1"},
		}

		GetUser(c)

		var user domain.User
		err := json.Unmarshal(response.Body.Bytes(), &user)
		assert.Nil(t, err)

	})
}
