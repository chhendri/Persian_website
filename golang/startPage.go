package main

import (
  "net/http"
  "html/template"
)

type buttonValue struct {
  Val string `json:"val"`
}


func startPageGet(w http.ResponseWriter, r *http.Request)  {
  tmpl := template.Must(template.ParseFiles("html_files/startPage.html"))
  tmpl.Execute(w, nil)
}


func startPagePost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  switch buttonVal.Val {
  case "Next":
    http.Redirect(w, r, "/listWords", http.StatusFound)
  case "wordTrainer":
    http.Redirect(w, r, "/wordTrain", http.StatusFound)
  case "imageTrainer":
    http.Redirect(w, r, "/imageTrain", http.StatusFound)
  }
}
