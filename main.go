package main

import (
  "io"
  "log"
  "net/http"
  "github.com/stealthrocket/net/wasip1"
)

func main() {
  listener, err := wasip1.Listen("tcp", "127.0.0.1:3000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "Hello world!\n")
  })
  server := &http.Server{
    Handler: handler,
  }

  if err := server.Serve(listener); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}

