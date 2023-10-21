package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
);

const PORT = ":8000";

func index(writer http.ResponseWriter, reader *http.Request) {
  template := template.Must(template.ParseFiles("index.html"));

  template.Execute(writer, nil);
}

func main() {
  fmt.Printf("Listening on http://localhost%s\n", PORT);

  http.HandleFunc("/", index);

  log.Fatal(http.ListenAndServe(PORT, nil));
}
