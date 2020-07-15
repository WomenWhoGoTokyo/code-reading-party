package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WomenWhoGoTokyo/code-reading-party/collections/collection"
)

type Handler struct {
	collection *collection.Collection
}

func NewHandlers(collection *collection.Collection) *Handler {
	return &Handler{
		collection: collection,
	}
}

func (h *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	list, err := h.collection.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: JSON でレスポンス
	fmt.Println(list)
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	var art collection.Art
	err = json.Unmarshal(b, &art)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	// TODO: Validation

	if err := h.collection.Add(&art); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: JSON でレスポンス
}
