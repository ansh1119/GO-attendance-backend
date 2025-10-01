# Attendance App

A backend system built with **Go** and **MongoDB** to manage attendance for bootcamps, workshops, and other events efficiently. The system allows organizers to mark attendance, fetch participant lists, and generate reports via RESTful APIs.

---

## Features

- **REST API**: Endpoints for marking attendance, fetching participants, and generating reports.
- **Database Management**: Stores event details, participants, and attendance maps per date using MongoDB.
- **Bulk Updates**: Supports bulk attendance updates with high throughput (~500+ entries/sec).
- **Concurrency Handling**: Leverages Go's goroutines and mutexes to safely handle multiple simultaneous updates.
- **Email Management**: Filters and manages participant emails during bulk uploads to prevent duplicates.
- **Scalable**: Designed to handle thousands of participants across multiple events.

---

## Tech Stack

- **Backend:** Go (Golang)
- **Database:** MongoDB
- **API:** RESTful JSON API
- **Concurrency:** Goroutines and Mutexes
- **Dependencies:** `go.mongodb.org/mongo-driver` (MongoDB driver)

---

## API Endpoints

### 1. Mark Attendance
