package middlewares

import (
	"net/http"

	"example.com/rapi/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(c *gin.Context){
	token:=c.Request.Header.Get("Authorisation")
   if token==""{
	   c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not Authorised"})
   }

	userId,err:=utils.VerifyToken(token)

	if err!=nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not Authorised"})
		return
	}
	c.Set("userId",userId)
   c.Next()
}