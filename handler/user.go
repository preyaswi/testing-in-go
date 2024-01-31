package handler

import (
	"net/http"

	"ks/domain"
	"ks/repository"

	"github.com/gin-gonic/gin"
)

func Hi(c *gin.Context) {
	c.JSON(http.StatusOK, "hello there")
}
func SignUp(c *gin.Context) {
	var user domain.SignUp
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	err := repository.GetByName(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creating user failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user created succesfuly"})
	}
}
func Login(c *gin.Context) {
	var input domain.SignUp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.Login(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user loged in succesfuly"})
	}

}

