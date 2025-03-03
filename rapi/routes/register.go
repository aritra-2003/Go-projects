package routes

import (
	"net/http"
	"strconv"

	"example.com/rapi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context){
	userId:=c.GetInt64("userId")
	eid := c.Param("id")
	st, err := strconv.ParseInt(eid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
   event,err:=models.GetEvent(st)
   if err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{
		"message":"can not register for an event"})
		return


   }
  err= event.Register(userId)
  if err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"message":"could not register for an event"})
	return
  }
  c.JSON(http.StatusCreated,gin.H{"message":"registered"})



}
func DeleteAnEvent(c *gin.Context){
	userId:=c.GetInt64("userId")
	eid := c.Param("id")
	st, err := strconv.ParseInt(eid, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"could not  find an event"})
	}
    var event models.Event
	event.ID=st
   err= event.CancelRegistration(userId)
   if err!=nil{
	 c.JSON(http.StatusInternalServerError,gin.H{"message":"could not cancel an event"})
	 return
   }
   c.JSON(http.StatusCreated,gin.H{"message":"event cancelled"})
 

}