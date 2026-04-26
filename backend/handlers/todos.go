package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"crud-app/services"
)

type TodoHandler struct {
	svc *services.TodoService
}

func New(svc *services.TodoService) *TodoHandler {
	return &TodoHandler{svc: svc}
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/todos")
	path = strings.TrimPrefix(path, "/")

	switch {
	case r.Method == http.MethodGet && path == "":
		h.list(w, r)
	case r.Method == http.MethodPost && path == "":
		h.create(w, r)
	case r.Method == http.MethodPut && path != "":
		h.update(w, r, path)
	case r.Method == http.MethodDelete && path != "":
		h.delete(w, r, path)
	default:
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
	}
}

func (h *TodoHandler) list(w http.ResponseWriter, r *http.Request) {
	todos, err := h.svc.GetAll()
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonError(w, "invalid body", http.StatusBadRequest)
		return
	}
	todo, err := h.svc.Create(body.Title)
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) update(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var body struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonError(w, "invalid body", http.StatusBadRequest)
		return
	}
	todo, err := h.svc.Update(id, body.Title, body.Completed)
	if err != nil {
		if err.Error() == "not found" {
			jsonError(w, "not found", http.StatusNotFound)
		} else {
			jsonError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) delete(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.svc.Delete(id); err != nil {
		if err.Error() == "not found" {
			jsonError(w, "not found", http.StatusNotFound)
		} else {
			jsonError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
