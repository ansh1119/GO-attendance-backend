package router

import (
	"github.com/ansh1119/GO-attendance-backend.git/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(eventHandler *handlers.EventHandler) *gin.Engine {
	r := gin.Default()

	// Versioning (good practice: /api/v1)
	api := r.Group("/api/v1")
	{
		// Events group
		events := api.Group("/events")
		{
			events.POST("", eventHandler.CreateEvent)                                  // POST /api/v1/events
			events.GET("/:eventId", eventHandler.GetEvent)                             // GET  /api/v1/events/:eventId
			events.POST("/:eventId/attendance/:date", eventHandler.MarkAttendance)     // POST /api/v1/events/:eventId/attendance/:date
			events.GET("/:eventId/attendance/:date", eventHandler.GetAttendanceByDate) // GET /api/v1/events/:eventId/attendance/:date
			events.POST("/:eventId/addUsers", eventHandler.UploadRegisteredUsersCSV)
		}
	}

	return r
}
