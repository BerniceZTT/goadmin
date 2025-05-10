package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BerniceZTT/goadmin/controllers"
	"github.com/BerniceZTT/goadmin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

type UserControllerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	mockService *MockUserService
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockService = new(MockUserService)
	userController := controllers.NewUserController(suite.mockService)

	suite.router = gin.Default()
	suite.router.GET("/users/:id", userController.GetUser)
	suite.router.POST("/users", userController.CreateUser)
}

func (suite *UserControllerTestSuite) TestGetUserSuccess() {
	expectedUser := &models.User{
		Model: gorm.Model{ID: 1},
		Name:  "Test User",
		Email: "test@example.com",
	}

	suite.mockService.On("GetUser", uint(1)).Return(expectedUser, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response models.User
	json.Unmarshal(w.Body.Bytes(), &response)
	suite.Equal(expectedUser.ID, response.ID)
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
