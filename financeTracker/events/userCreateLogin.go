package events

import (
	"database/sql"
	
	"net/http"

	"example.com/financetracker/models"

	"example.com/financetracker/database"
	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context){
	var user models.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"invalid request"})
		return
	}
	query:=`INSERT INTO "users" (email,password) values($1, $2) RETURNING id`
	row:=database.DB.QueryRow(query,user.Email,user.Password)
	if err:=row.Scan(&user.ID);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)



}
func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var storedPassword string
	query:=`SELECT id,Password FROM "users" WHERE Email=$1`
	err := database.DB.QueryRow(query, user.Email).Scan(&user.ID, &storedPassword)
	
	if err == sql.ErrNoRows || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
    
	c.JSON(http.StatusOK, gin.H{"message": "login successful", "user_id": user.ID})
	

}