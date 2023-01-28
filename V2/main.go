package main

import (
  "html/template"
)

var templates *template.Template



func main() {
  connectDB()

  templates = template.Must(templates.ParseGlob("static/html_files/*.html"))

  Mux()
}
