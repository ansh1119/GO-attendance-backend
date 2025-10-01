package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/ansh1119/GO-attendance-backend.git/repository"

	"github.com/ansh1119/GO-attendance-backend.git/models"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	repo *repository.EventRepository
}

// constructor
func NewEventHandler(repo *repository.EventRepository) *EventHandler {
	return &EventHandler{repo: repo}
}

// Create a new event
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req models.Event
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ensure attendance map is not nil
	if req.Attendance == nil {
		req.Attendance = make(map[string][]string)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.repo.CreateEvent(ctx, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created"})
}

// Get event by ID
func (h *EventHandler) GetEvent(c *gin.Context) {
	eventID := c.Param("eventId")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event, err := h.repo.GetEventByEventID(ctx, eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// Mark attendance
func (h *EventHandler) MarkAttendance(c *gin.Context) {
	eventID := c.Param("eventId")
	date := c.Param("date")

	var body struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.repo.MarkAttendance(ctx, eventID, date, body.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "attendance marked"})
}

// Get attendance for a specific date
func (h *EventHandler) GetAttendanceByDate(c *gin.Context) {
	eventID := c.Param("eventId")
	date := c.Param("date")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	attendees, err := h.repo.GetAttendanceByDate(ctx, eventID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"eventId":   eventID,
		"date":      date,
		"attendees": attendees,
	})
}

func (h *EventHandler) UploadRegisteredUsersCSV(c *gin.Context) {
	// Get eventId from URL path
	eventID := c.Param("eventId")

	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file is required"})
		return
	}

	// Save temporarily
	filePath := "./tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create context with timeout for Mongo operation
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Call repository method
	if err := h.repo.AddRegisteredUsersFromCSV(ctx, eventID, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered users updated successfully"})
}
