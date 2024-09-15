package handlers

import (
	"cab-management/database"
	"cab-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterCab(c *gin.Context) {
	var newCab models.Cab
	if err := c.ShouldBindJSON(&newCab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCab.LastIdleTime = time.Now()
	database.DB.Create(&newCab)
	c.JSON(http.StatusOK, newCab)
}

func AddCity(c *gin.Context) {
	var newCity models.City
	if err := c.ShouldBindJSON(&newCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newCity)
	c.JSON(http.StatusOK, newCity)
}

func ChangeCabLocation(c *gin.Context) {
	var input struct {
		CabID  uint `json:"cab_id"`
		CityID uint `json:"city_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cab models.Cab
	database.DB.First(&cab, input.CabID)
	cab.CurrentCityID = input.CityID
	database.DB.Save(&cab)

	c.JSON(http.StatusOK, cab)
}

func ChangeCabState(c *gin.Context) {
	var input struct {
		CabID uint   `json:"cab_id"`
		State string `json:"state"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cab models.Cab
	database.DB.First(&cab, input.CabID)

	// Update the state and log history
	cab.State = input.State
	if input.State == "IDLE" {
		cab.LastIdleTime = time.Now()
	}

	database.DB.Save(&cab)

	cabHistory := models.CabHistory{
		CabID:     input.CabID,
		CityID:    cab.CurrentCityID,
		State:     input.State,
		Timestamp: time.Now(),
	}
	database.DB.Create(&cabHistory)

	c.JSON(http.StatusOK, cab)
}

func BookCab(c *gin.Context) {
	var input struct {
		CityID uint `json:"city_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cabs []models.Cab
	database.DB.Where("current_city_id = ? AND state = ?", input.CityID, "IDLE").Find(&cabs)

	if len(cabs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No cabs available"})
		return
	}

	var selectedCab models.Cab
	var maxIdleTime time.Duration
	for _, cab := range cabs {
		idleTime := time.Since(cab.LastIdleTime)
		if idleTime > maxIdleTime {
			maxIdleTime = idleTime
			selectedCab = cab
		}
	}

	// If there's a clash on idle time, randomly pick one
	if selectedCab.ID == 0 && len(cabs) > 0 {
		selectedCab = cabs[0]
	}

	// Assign cab to trip
	selectedCab.State = "ON_TRIP"
	database.DB.Save(&selectedCab)

	c.JSON(http.StatusOK, selectedCab)
}

func CabIdleTime(c *gin.Context) {
	var input struct {
		CabID uint `json:"cab_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cab models.Cab
	database.DB.First(&cab, input.CabID)

	idleDuration := time.Since(cab.LastIdleTime)
	c.JSON(http.StatusOK, gin.H{"idle_time": idleDuration})
}

func GetCabHistory(c *gin.Context) {
	var input struct {
		CabID uint `json:"cab_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var history []models.CabHistory
	database.DB.Where("cab_id = ?", input.CabID).Find(&history)

	c.JSON(http.StatusOK, history)
}
