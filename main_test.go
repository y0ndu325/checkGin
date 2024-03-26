package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleRequest(w *httptest.ResponseRecorder, r *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, r)
}

func TestAlbumList(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	handleRequest(w, req)
	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestAlbum(t *testing.T) {
	albumID := "1"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/"+albumID, nil)
	handleRequest(w, req)
	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestAlbumNotFound(t *testing.T) {
	albumID := "11"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/"+albumID, nil)
	handleRequest(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatal("status => 404")
	}
}

func TestDeleteAlbum(t *testing.T) {
	albumId := "1"
	request, _ := http.NewRequest("DELETE", "/albums/"+albumId, nil)
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status => 204")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	albumID := "11"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/"+albumID, nil)
	handleRequest(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatal("status => 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	albumID := "2"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/"+albumID, strings.NewReader(`{"title":"test"}`))
	handleRequest(w, req)
	if w.Code != http.StatusOK {
		t.Fatal("status => ok")
	}
}

func TestCreateAlbumBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	handleRequest(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatal("status => 400")
	}
}

func TestCreateAlbum(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201")
	}
}
