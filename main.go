package main;

import (
	"fmt"
	"io"
	"log"
	"net/http"
);

const PORT = ":8000";

func index(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<h1>Hello, world!</h1>")
}

func main() {
  fmt.Printf("Listening on http://localhost%s\n", PORT);

  http.HandleFunc("/", index);

  log.Fatal(http.ListenAndServe(PORT, nil));
}
