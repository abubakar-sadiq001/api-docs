package handlers

import (
	"log"
	"net/http"

	"github.com/abubakar-sadiq001/model"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

type errResponse struct {
	Message string
}

// GetUsers return list of all users from the database
// @Summary return list of all
// @Description return list of all users from the database
// @Tags Users
// @Security bearerToken
// @Success 200 {object} userResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {

	users := model.ListUsers()

	c.JSON(http.StatusOK, userResponse{Data: users})
}

// CreateUser create new user
// @Summary create new user
// @Description create new user
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} errResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {

	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := model.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	auth := c.GetHeader("Authorization")
	log.Println("Authorization", auth)

	c.JSON(http.StatusCreated, req)
}
