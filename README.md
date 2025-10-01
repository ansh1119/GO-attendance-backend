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

POST /attendance

**Request Body:**
{
  "eventId": "E123",
  "date": "2025-09-01",
  "participantEmails": ["user1@example.com", "user2@example.com"]
}


**Response:**

{
  "success": true,
  "message": "Attendance marked successfully"
}

**2. Fetch Participants**
GET /participants?eventId=E123&date=2025-09-01


**Response:**

{
  "participants": ["user1@example.com", "user2@example.com"]
}

**3. Generate Attendance Report**
GET /report?eventId=E123


Response:

{
  "eventId": "E123",
  "attendance": {
    "2025-09-01": ["user1@example.com", "user2@example.com"],
    "2025-09-02": ["user3@example.com"]
  }
}

**Installation**

**Clone the repository:**

git clone https://github.com/ansh1119/GO-attendance-backend.git
cd GO-attendance-backend


Install dependencies: go mod tidy


Run the server: go run main.go
