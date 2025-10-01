package models

import "time"

// Event represents an event document in MongoDB
type Event struct {
	ID              string              `bson:"_id,omitempty" json:"id"`
	Name            string              `bson:"name" json:"name"`
	Attendance      map[string][]string `bson:"attendance" json:"attendance"`
	RegisteredUsers []string            `bson:"registeredUsers" json:"registeredUsers"`
	StartDate       time.Time           `bson:"startDate" json:"startDate"`
	EndDate         time.Time           `bson:"endDate" json:"endDate"`
}
