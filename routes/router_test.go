package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/raviwu/gobookstore/models"
)

var (
	db *gorm.DB
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestFindBooks(t *testing.T) {
	db = models.SetupModels()
	defer db.Close()

	body := `{"data":[]}`

	router := SetupRouter(db)

	w := performRequest(router, "GET", "/books")

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

	cleaner := deleteCreatedEntities(db)
	defer cleaner()

	db.Create(&models.Book{Title: "book", Author: "ravi"})

	var books []models.Book
	db.Find(&books)

	t.Log(books)

	body := fmt.Sprintf(`{"data":[{"id":%v,"title":"book","author":"ravi"}]}`, books[0].ID)

	router := SetupRouter(db)
	w := performRequest(router, "GET", "/books")

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}

func deleteCreatedEntities(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})

	return func() {
		defer db.Callback().Create().Remove(hookName)
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}

		for _, entry := range entries {
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}
