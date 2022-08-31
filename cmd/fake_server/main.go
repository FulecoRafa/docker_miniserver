package main

import (
  "io"
  "log"
  "net/http"
  "os"
)

var httplog *log.Logger = log.New(
  os.Stdout,
  "Endpoint hit: ",
  log.Ldate|log.LstdFlags|log.Lshortfile,
)

func handleRedirect(w http.ResponseWriter, r *http.Request) {
  httplog.Println("handleRedirect")
  // Make get request to the redirect url
  resp, err := http.Get(os.Getenv("REDIRECT_URL"))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  // Copy the response body to the response writer
  io.WriteString(w, "Redirected!\n")
  io.Copy(w, resp.Body)

  resp, err = http.Get(os.Getenv("DATABASE_URL") + "/travel")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  io.Copy(w, resp.Body)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  httplog.Println("handleRequest")
  resp, err := http.Get(os.Getenv("DATABASE_URL") + "/user")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  io.WriteString(w, "Hit last point!\n")
  io.Copy(w, resp.Body)
}

func main() {
  log.Println("Starting server...")
  http.HandleFunc("/", handleRequest)
  http.HandleFunc("/redirect", handleRedirect)
  http.ListenAndServe(":4444", nil)
}
