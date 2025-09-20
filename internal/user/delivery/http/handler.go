package http

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	uc     *usecase.UserUseCase
	logger *logrus.Logger
}

func NewUserHandler(router *gin.Engine, uc *usecase.UserUseCase, logger *logrus.Logger) {
	handler := &UserHandler{uc: uc, logger: logger}
	router.GET("/users", handler.GetUsers)    // covers both list and filter by name
	router.GET("/users/:id", handler.GetUser) // find by ID
	router.GET("/users/search/:name", handler.SearchUsers)
	router.POST("/users", handler.CreateUser)
}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve all users
// @Tags users
// @Produce json
// @Success 200 {array} entity.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// SearchUsers godoc
// @Summary Search users by name
// @Description Retrieve users filtered by name
// @Tags users
// @Produce json
// @Param name query string true "User name"
// @Success 200 {array} entity.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/search [get]
func (h *UserHandler) SearchUsers(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing name query param"})
		return
	}

	users, err := h.uc.GetUsersByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no users found: " + name})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.User true "User info"
// @Success 201 {object} entity.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.uc.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
