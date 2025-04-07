
# 📝 To-Do List API - Documentation

This is a simple RESTful backend API built with **Go (Golang)** for managing a list of to-do items. The API allows clients to create, read, update, and delete (CRUD) to-do tasks.

---

## 🌐 Base URL
```
http://localhost:8080
```

---

## 📌 Available Endpoints

### 1. `GET /todos`
#### Description:
Retrieve all to-do items.

#### Request:
- **Method:** GET  
- **Authentication:** Not required  
- **Headers:** None

#### Response:
- **Status:** `200 OK`  
- **Body Example:**
```json
[
  {
    "id": 1,
    "title": "Learn Golang",
    "note": "Study HTTP package and routing",
    "completed": false
  }
]
```

---

### 2. `GET /todos/{id}`
#### Description:
Retrieve a specific to-do item by its ID.

#### Request:
- **Method:** GET  
- **Path Parameter:** `id` (integer)  
- **Example:** `/todos/1`

#### Response:
- **Status:** `200 OK`  
```json
{
  "id": 1,
  "title": "Learn Golang",
  "note": "Study HTTP package and routing",
  "completed": false
}
```

- **Error Response:** `404 Not Found` if the ID is not found

---

### 3. `POST /todos`
#### Description:
Create a new to-do item.

#### Request:
- **Method:** POST  
- **Headers:**
  - `Content-Type: application/json`  
- **Body Example:**
```json
{
  "title": "Read documentation",
  "note": "Explore encoding/json in Go",
  "completed": false
}
```

#### Response:
- **Status:** `201 Created`  
- **Body Example:**
```json
{
  "id": 2,
  "title": "Read documentation",
  "note": "Explore encoding/json in Go",
  "completed": false
}
```

- **Error Response:** `400 Bad Request` if the request body is invalid JSON

---

### 4. `PUT /todos/{id}`
#### Description:
Update an existing to-do item by ID.

#### Request:
- **Method:** PUT  
- **Path Parameter:** `id` (integer)  
- **Headers:**
  - `Content-Type: application/json`  
- **Body Example:**
```json
{
  "title": "Update title",
  "note": "Update the note content",
  "completed": true
}
```

#### Response:
- **Status:** `204 No Content`

- **Error Responses:**
  - `400 Bad Request` if the request body is invalid JSON  
  - `404 Not Found` if the item with the specified ID does not exist

---

### 5. `DELETE /todos/{id}`
#### Description:
Delete a to-do item by ID.

#### Request:
- **Method:** DELETE  
- **Path Parameter:** `id` (integer)

#### Response:
- **Status:** `204 No Content`

- **Error Response:** `404 Not Found` if the ID is not found

---

## ⚠️ Error Handling

| Status Code | Description                    |
|-------------|--------------------------------|
| 400         | Bad Request (Invalid input)    |
| 404         | Not Found (Item not found)     |
| 405         | Method Not Allowed             |

---

## 📦 Project Structure

```
todo-api/
│
├── main.go               # Entry point of the app
├── handler/              # HTTP request handlers
│   └── todo.go
├── model/                # To-do data model
│   └── todo.go
├── storage/              # In-memory data store logic
│   └── memory.go
├── go.mod                # Go module file
└── README.md             # API documentation (this file)
```

---

## 🔧 Technologies Used
- Go (Golang)
- net/http (standard Go HTTP package)
- JSON for data exchange (via `encoding/json`)

---

## ⚙️ Notes
- This API uses in-memory data storage only (not persistent).
- All data will be lost when the server is restarted.
- No authentication is required for this demo API.

---

## ▶️ How to Run

1. Make sure you have Go installed (`go version`)
2. Clone the repo or place the files in a folder:
```bash
go mod init todo-api
go run main.go
```
3. The server will start at: `http://localhost:8080`

---

## 🧪 Testing

You can test the API using tools like:
- [Postman](https://www.postman.com/)
- [curl](https://curl.se/)
- [Insomnia](https://insomnia.rest/)

---

## ✍️ Author

This project was created as part of a backend programming assignment using Golang.
