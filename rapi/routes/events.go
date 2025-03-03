package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/rapi/models"
	
	"github.com/gin-gonic/gin"
)

// Handler to get all events
func getEvents(c *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Handler to get a single event
func getEvent(c *gin.Context) {
	eid := c.Param("id")
	st, err := strconv.ParseInt(eid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEvent(st)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Handler to create an event
func createEvent(c *gin.Context) {
	var event models.Event
   
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    userId:= c.GetInt64("userId")
	event.ID = userId
	event.UserId = 1
	event.Date = time.Now()

	err := event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": event})
}

// Handler to update an event
func updateEvent(c *gin.Context) {
	eid := c.Param("id")
	st, err := strconv.ParseInt(eid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
    userId:= c.GetInt64("userId")
	event, err := models.GetEvent(st)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	if event.UserId!=userId{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"not autthorised"})

	}

	var updateEvent models.Event
	if err := c.ShouldBindJSON(&updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = st
	err = models.UpdateEvent(&updateEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
func deleteEvent(c* gin.Context){
	eid := c.Param("id")
	st, err := strconv.ParseInt(eid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
    userId:= c.GetInt64("userId")
  event, err := models.GetEvent(st)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	if event.UserId!=userId{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"not autthorised to delate"})

	}
    err=event.Delete()
	if err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
	

}
