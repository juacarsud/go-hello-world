package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandlerFunc(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    w := httptest.NewRecorder()
    handlerFunc(w, req)
    res := w.Result()
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    if string(data) == "Hello World, server is listening on 9090" {
        t.Errorf("expected Hello Wold got %v", string(data))
    }
}