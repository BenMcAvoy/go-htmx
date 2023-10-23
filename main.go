package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql" // or postgres, sqlite3, etc.
	"github.com/joho/godotenv"
);

type Film struct {
  Title    string;
  Director string;
}

const PORT = ":8000";

func addFilm(writer http.ResponseWriter, reader *http.Request) {
	title    := reader.PostFormValue("title");
	director := reader.PostFormValue("director");

  // If either of the fields are empty, do nothing.
  if len(title) == 0 || len(director) == 0 {
    return;
  }

  // Simulate expensive task.
  time.Sleep(1 * time.Second);

  log.Printf("Adding film:\n                     └ %s - %s", title, director);

	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director);
  tmpl, _ := template.New("t").Parse(htmlStr);

  tmpl.Execute(writer, nil);
}

func index(writer http.ResponseWriter, reader *http.Request) {
  username := os.Getenv("USERNAME");
  password := os.Getenv("PASSWORD");
  ip := os.Getenv("IP");
  port := os.Getenv("PORT");
  database := os.Getenv("DATABASE");

  connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, ip, port, database);

  db, err := sql.Open("mysql", connectUrl);

  if err != nil {
    log.Fatalf("could not connect to database: %v", err);
  }

  log.Println("Connected to database.");

  template := template.Must(template.ParseFiles("static/index.html"));

  rows, err := db.Query("SELECT Title, Director FROM Films");

  if err != nil {
    log.Fatal(err);
  }
  
  defer rows.Close();

  films := make(map[string][]Film);

  for rows.Next() {
    var film Film;
    
    err = rows.Scan(&film.Title, &film.Director);

    if err != nil {
      log.Fatal(err);
    }

    // Append the film to the slice
    films["Films"] = append(films["Films"], film);
  }

  err = rows.Err();

  if err != nil {
    log.Fatal(err);
  }

  template.Execute(writer, films);
}

func main() {
  if godotenv.Load() != nil {
    log.Fatal("Failed to load .env");
  }

  log.Printf("Listening on http://localhost%s\n", PORT);

  http.HandleFunc("/add-film/", addFilm);
  http.HandleFunc("/", index);

  fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

  log.Fatal(http.ListenAndServe(PORT, nil));
}
