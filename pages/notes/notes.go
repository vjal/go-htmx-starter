package notes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"goth/pages"
)

type Note = struct {
	Title string
	Body  string
	Id    int
}

var notes = []Note{
	{"Note 1", "This is the first note", 0},
	{"Note 2", "This is the second note", 1},
	{"Note 3", "This is the third note", 2},
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	c := NotesPage(notes)
	err := pages.Layout(c, "Gonotes").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func NoteHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid note id", http.StatusBadRequest)
		return
	}
	c := NotePage(notes[id])
	err = pages.Layout(c, "Gonotes").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
