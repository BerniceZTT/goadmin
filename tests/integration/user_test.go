package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BerniceZTT/goadmin/config"
	"github.com/BerniceZTT/goadmin/models"
	"github.com/BerniceZTT/goadmin/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type IntegrationTestSuite struct {
	suite.Suite
	db     *gorm.DB
	router *gin.Engine
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// 初始化测试数据库
	os.Setenv("DB_NAME", "test_db")
	config.ConnectDB()
	suite.db = config.DB
	suite.db.AutoMigrate(&models.User{})

	// 初始化路由
	suite.router = routes.SetupRouter()
}

func (suite *IntegrationTestSuite) TearDownTest() {
	// 清空测试数据
	suite.db.Exec("DELETE FROM users")
}

func (suite *IntegrationTestSuite) TestUserCRUD() {
	// 创建用户
	newUser := models.User{
		Name:  "Integration Test",
		Email: "integration@test.com",
	}

	userJSON, _ := json.Marshal(newUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var createdUser models.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)

	// 查询用户
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/users/"+fmt.Sprint(createdUser.ID), nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var fetchedUser models.User
	json.Unmarshal(w.Body.Bytes(), &fetchedUser)
	suite.Equal(newUser.Name, fetchedUser.Name)
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
