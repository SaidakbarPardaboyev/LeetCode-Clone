package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	filter := model.UserFilter{}
	fullName, hasKey := c.GetQuery("full_name")
	if hasKey {
		filter.FullName = &fullName
	}
	username, hasKey := c.GetQuery("username")
	if hasKey {
		filter.FullName = &username
	}

	users, err := h.UserRepo.GetUsers(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Error getting users",
		})
		log.Println("Error getting users", err)
		return
	}
	c.JSON(http.StatusOK, users)

}

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserRepo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting user by Id",
		})
		log.Println("Error getting user by Id", err)
		return
	}
	c.JSON(200, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	newUser := model.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error decoding user",
		})
		log.Println("Error while decoding user", err)
		return
	}
	err = h.UserRepo.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating user",
		})
		log.Println("Error while creating user", err)
		return
	}
}

func (h *Handler) UpdateUser(c *gin.Context) {
	user := model.User{}

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error decoding user",
		})
		log.Println("Error while decoding user", err)
		return
	}

	user.Id = c.Param("id")
	err = h.UserRepo.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating use",
		})
		log.Println("Error while updating user", err)
		return
	}
}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	err := h.UserRepo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating use",
		})
		log.Println("Error while deleting user", err)
		return
	}
}
