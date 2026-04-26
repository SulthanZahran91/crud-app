package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"crud-app/handlers"
	"crud-app/repository"
	"crud-app/services"
)

func setup(t *testing.T) *handlers.TodoHandler {
	t.Helper()
	f, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	t.Cleanup(func() { os.Remove(f.Name()) })

	repo, err := repository.New(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	return handlers.New(services.New(repo))
}

func TestListEmpty(t *testing.T) {
	h := setup(t)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200 got %d", w.Code)
	}
	var todos []map[string]interface{}
	json.NewDecoder(w.Body).Decode(&todos)
	if len(todos) != 0 {
		t.Fatalf("want empty list got %v", todos)
	}
}

func TestCreateAndList(t *testing.T) {
	h := setup(t)

	body := `{"title":"buy milk"}`
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201 got %d: %s", w.Code, w.Body)
	}

	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	var todos []map[string]interface{}
	json.NewDecoder(w.Body).Decode(&todos)
	if len(todos) != 1 || todos[0]["title"] != "buy milk" {
		t.Fatalf("unexpected todos: %v", todos)
	}
}

func TestUpdate(t *testing.T) {
	h := setup(t)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(`{"title":"test"}`))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	var created map[string]interface{}
	json.NewDecoder(w.Body).Decode(&created)

	id := int(created["id"].(float64))
	body, _ := json.Marshal(map[string]interface{}{"title": "updated", "completed": true})
	req = httptest.NewRequest(http.MethodPut, "/todos/"+itoa(id), bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200 got %d: %s", w.Code, w.Body)
	}
}

func TestDelete(t *testing.T) {
	h := setup(t)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(`{"title":"delete me"}`))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	var created map[string]interface{}
	json.NewDecoder(w.Body).Decode(&created)

	id := int(created["id"].(float64))
	req = httptest.NewRequest(http.MethodDelete, "/todos/"+itoa(id), nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("want 204 got %d", w.Code)
	}
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
