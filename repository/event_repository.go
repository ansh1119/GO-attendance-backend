package repository

import (
	"context"
	"errors"

	"github.com/ansh1119/GO-attendance-backend.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	collection *mongo.Collection
}

// Initialize repository with Mongo collection
func NewEventRepository(db *mongo.Database, collectionName string) *EventRepository {
	return &EventRepository{
		collection: db.Collection(collectionName),
	}
}

// Create a new event
func (r *EventRepository) CreateEvent(ctx context.Context, e *models.Event) error {
	if e.Attendance == nil {
		e.Attendance = make(map[string][]string)
	}
	_, err := r.collection.InsertOne(ctx, e)
	return err
}

// Get event by EventID
func (r *EventRepository) GetEventByEventID(ctx context.Context, eventID string) (*models.Event, error) {
	var event models.Event
	objID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return nil, errors.New("invalid event ID")
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}
	return &event, nil
}

// Mark attendance: add email to attendance[date] if not already present
func (r *EventRepository) MarkAttendance(ctx context.Context, eventID, date, email string) error {
	objID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return errors.New("invalid event ID")
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$addToSet": bson.M{
			"attendance." + date: email,
		},
	}
	_, er := r.collection.UpdateOne(ctx, filter, update)
	return er
}

// Get attendance for a specific date
func (r *EventRepository) GetAttendanceByDate(ctx context.Context, eventID, date string) ([]string, error) {
	event, err := r.GetEventByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	attendees := event.Attendance[date]
	if attendees == nil {
		attendees = []string{}
	}
	return attendees, nil
}
