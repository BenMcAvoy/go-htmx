package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
);

type Film struct {
  Title    string;
  Director string;
}

const PORT = ":8000";

func index(writer http.ResponseWriter, reader *http.Request) {
  template := template.Must(template.ParseFiles("index.html"));

  films := map[string][]Film {
    "Films": {
      {Title: "The Godfather", Director: "Fancis Ford Coppola"},
      {Title: "Blade Runner", Director: "Ridley Scott"},
      {Title: "The Thing", Director: "John Carpenter"},
    },
  }

  template.Execute(writer, films);
}

func main() {
  fmt.Printf("Listening on http://localhost%s\n", PORT);

  http.HandleFunc("/", index);

  log.Fatal(http.ListenAndServe(PORT, nil));
}
