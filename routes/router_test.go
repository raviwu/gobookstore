package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/raviwu/gobookstore/models"
)

var (
	db *gorm.DB
)

func performRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, ioutil.NopCloser(strings.NewReader(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestFindBooks(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()
	db.Exec(`DELETE FROM books`)

	body := `{"data":[]}`

	router := SetupRouter(db)

	w := performRequest(router, "GET", "/books", "")

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}

func TestCreateBook(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()
	db.Exec(`DELETE FROM books`)

	router := SetupRouter(db)

	w := performRequest(router, "POST", "/books", `{"title":"new book","author":"another ravi"}`)

	var books []models.Book
	db.Find(&books)

	body := fmt.Sprintf(`{"data":{"id":%v,"title":"new book","author":"another ravi"}}`, books[0].ID)

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}

func TestFindBook(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()
	db.Exec(`DELETE FROM books`)

	book := models.Book{Title: "new new book", Author: "hey ravi"}
	db.Create(&book)

	var books []models.Book
	db.Find(&books)

	body := fmt.Sprintf(`{"data":{"id":%v,"title":"%s","author":"%s"}}`, books[0].ID, books[0].Title, books[0].Author)
	path := fmt.Sprintf(`/books/%v`, books[0].ID)

	router := SetupRouter(db)
	w := performRequest(router, "GET", path, "")

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}

func TestUpdateBook(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()
	db.Exec(`DELETE FROM books`)

	book := models.Book{Title: "new new book", Author: "hey ravi"}
	db.Create(&book)

	body := fmt.Sprintf(`{"data":{"id":%v,"title":"%s","author":"%s"}}`, book.ID, "aaa", "bbb")

	data := fmt.Sprintf(`{"title":"%s","author":"%s"}`, "aaa", "bbb")
	path := fmt.Sprintf(`/books/%v`, book.ID)

	router := SetupRouter(db)
	w := performRequest(router, "PATCH", path, data)

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}

func TestDeleteBook(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()
	db.Exec(`DELETE FROM books`)

	book := models.Book{Title: "new new book", Author: "hey ravi"}
	db.Create(&book)

	path := fmt.Sprintf(`/books/%v`, book.ID)

	router := SetupRouter(db)
	w := performRequest(router, "DELETE", path, "")

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	var books []models.Book
	db.Find(&books)

	if len(books) > 0 {
		t.Errorf("expected: %s\nactual: %v", "[]", books)
	}
}
