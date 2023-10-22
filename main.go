package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
);

type Film struct {
  Title    string;
  Director string;
}

const PORT = ":8000";

func addFilm(writer http.ResponseWriter, reader *http.Request) {
  time.Sleep(1 * time.Second);

	title    := reader.PostFormValue("title");
	director := reader.PostFormValue("director");

  // If either of the fields are empty, do nothing.
  if len(title) == 0 || len(director) == 0 {
    return;
  }

  fmt.Printf("Adding film:\n  %s - %s", title, director);

	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director);
  tmpl, _ := template.New("t").Parse(htmlStr);

  tmpl.Execute(writer, nil);
}

func index(writer http.ResponseWriter, reader *http.Request) {
  fmt.Println("Returning main page.");

  template := template.Must(template.ParseFiles("static/index.html"));

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

  http.HandleFunc("/add-film/", addFilm);
  http.HandleFunc("/", index);

  fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

  log.Fatal(http.ListenAndServe(PORT, nil));
}
