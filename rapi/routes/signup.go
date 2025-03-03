package routes

import "github.com/gin-gonic/gin"
import "net/http"
import "example.com/rapi/models"
import "example.com/rapi/utils"
func signup(c *gin.Context) {
	// Implement the signup handler
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}
func login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"messgae":"User CANT LOGIN"})
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Could Not Authenticate"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful","token":token})


}
