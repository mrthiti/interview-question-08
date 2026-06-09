package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"thaibev-assignment/backend/database"
	"thaibev-assignment/backend/handlers"
	"thaibev-assignment/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&models.Question{})
	database.DB = db
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/questions", handlers.GetQuestions)
	r.POST("/api/questions", handlers.CreateQuestion)
	r.DELETE("/api/questions/:id", handlers.DeleteQuestion)
	return r
}

func TestGetQuestions_Empty(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/questions", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String())
}

func TestCreateQuestion(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{
		"question": "What is Go?",
		"choice_1": "A language",
		"choice_2": "A framework",
		"choice_3": "A database",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/questions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, "What is Go?", result["question"])
	assert.Equal(t, float64(1), result["order_number"])
}

func TestCreateQuestion_MissingQuestion(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{"choice_1": "only option"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/questions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteQuestion_RenumbersRemaining(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	// create 3 questions
	for _, q := range []string{"Q1", "Q2", "Q3"} {
		body, _ := json.Marshal(map[string]string{"question": q})
		req, _ := http.NewRequest("POST", "/api/questions", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(httptest.NewRecorder(), req)
	}

	// get question list to find IDs
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/questions", nil)
	r.ServeHTTP(w, req)

	var questions []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &questions)
	assert.Equal(t, 3, len(questions))

	// delete Q2 (index 1)
	q2ID := int(questions[1]["ID"].(float64))
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("DELETE", "/api/questions/"+strconv.Itoa(q2ID), nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)

	// verify remaining questions are re-numbered 1, 2
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/api/questions", nil)
	r.ServeHTTP(w3, req3)

	var remaining []map[string]interface{}
	json.Unmarshal(w3.Body.Bytes(), &remaining)
	assert.Equal(t, 2, len(remaining))
	assert.Equal(t, float64(1), remaining[0]["order_number"])
	assert.Equal(t, "Q1", remaining[0]["question"])
	assert.Equal(t, float64(2), remaining[1]["order_number"])
	assert.Equal(t, "Q3", remaining[1]["question"])
}

func TestDeleteQuestion_NotFound(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/questions/999", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
