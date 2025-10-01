package repository

import (
	"context"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *EventRepository) AddRegisteredUsersFromCSV(ctx context.Context, eventID, filePath string) error {
	// Open CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all rows
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil // empty file
	}

	// Detect email column in the first row
	emailColIndex := -1
	for i, cell := range rows[0] {
		if strings.EqualFold(strings.TrimSpace(cell), "email") || strings.EqualFold(strings.TrimSpace(cell), "mail") {
			emailColIndex = i
			break
		}
	}
	if emailColIndex == -1 {
		return nil // no email column found
	}

	// Collect emails concurrently
	var emails []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, row := range rows[1:] { // skip header
		if emailColIndex >= len(row) {
			continue
		}
		email := strings.TrimSpace(row[emailColIndex])
		if email == "" {
			continue
		}

		wg.Add(1)
		go func(e string) {
			defer wg.Done()
			mu.Lock()
			emails = append(emails, e)
			mu.Unlock()
		}(email)
	}

	wg.Wait()

	objID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return errors.New("invalid event ID")
	}

	// Update MongoDB using $addToSet to avoid duplicates
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$addToSet": bson.M{
			"registeredUsers": bson.M{"$each": emails},
		},
	}

	_, err = r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(false))
	if err != nil {
		log.Println("Mongo update error:", err)
	}
	return err
}
