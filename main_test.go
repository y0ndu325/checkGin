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

func createTestAlbum() album {
	TestAlbum := album{ID: "2", Title: "test", Artist: "test", Price: 123.0}
	storage.Create(TestAlbum)
	return TestAlbum
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
	testAlbum := createTestAlbum()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/"+testAlbum.ID, nil)
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
	TestAlbum := createTestAlbum()
	request, _ := http.NewRequest("DELETE", "/albums/"+TestAlbum.ID, nil)
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status => 204")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	albumID := "11"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/"+albumID, strings.NewReader(`{"title":"test"}`))
	handleRequest(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatal("status => 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	TestAlbum := createTestAlbum()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/"+TestAlbum.ID, strings.NewReader(`{"title":"test"}`))
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
