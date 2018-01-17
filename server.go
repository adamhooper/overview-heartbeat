package main

import (
  "log"
  "io"
  "net/http"
)

var showHtml = "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\"><title>Null Plugin</title></head><body><h1>Null Plugin</h1><p>This plugin does nothing. How did you find it?</p></body></html>"

func getShow(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Header().Set("Cache-Control", "public, max-age=3600")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, showHtml)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.Header().Set("Cache-Control", "public, max-age=3600")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, "OK")
}

func getMetadata(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Cache-Control", "public, max-age=3600")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, "{}")
}

func main() {
  http.HandleFunc("/", getRoot)
  http.HandleFunc("/show", getShow)
  http.HandleFunc("/metadata", getMetadata)
  err := http.ListenAndServe(":80", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
