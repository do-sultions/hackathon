package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  // Define some data to pass to the template
  data := struct {
    Title string
  }{
    Title: "My Go Web App",
  }

  // Parse the template file
  tmpl, err := template.ParseFiles("templates/home.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // Execute the template with data
  err = tmpl.Execute(w, data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8080", nil)
}