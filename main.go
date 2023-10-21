package main;

import (
	"fmt"
	"log"
	"net/http"
);

const PORT = ":8000";

func main() {
  fmt.Printf("Listening on http://localhost%s\n", PORT);
  log.Fatal(http.ListenAndServe(PORT, nil));
}
