package main

import (
  "net/http"
  "html/template"
)

type buttonValue struct {
  Val string `json:"val"`
}

var Language_user string
var Language_to_learn string

// Display of the Start page to choose the language to train and the method
func startPageGet(w http.ResponseWriter, r *http.Request)  {
  tmpl := template.Must(template.ParseFiles("html_files/startPage.html"))
  tmpl.Execute(w, nil)
}

// Choose only the method, with default language to learn is Farsi
func startPagePost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  switch buttonVal.Val {
  case "wordList":
    http.Redirect(w, r, "/listWords", http.StatusFound)
  case "wordTrainer":
    http.Redirect(w, r, "/wordTrain", http.StatusFound)
  case "imageTrainer":
    http.Redirect(w, r, "/imageTrain", http.StatusFound)
  }
}

// Choose the language and handle the method
func startPageLangPost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  switch buttonVal.Val {
  // Handle language
  case "farsi":
    Language_user = "farsi"
    tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_farsi.html"))
    tmpl.Execute(w, nil)
  case "french":
    Language_user = "french"
    tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_french.html"))
    tmpl.Execute(w, nil)
  case "german":
    Language_user = "german"
    tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_german.html"))
    tmpl.Execute(w, nil)
  }
}


// Choose the language to learn and handle the method
func startPageLangToLearnPost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  // Define the language to learn
  switch buttonVal.Val {
  case "farsi":
    Language_to_learn = "farsi"
  case "french":
    Language_to_learn = "french"
  case "german":
    Language_to_learn = "german"
  }

  // Define the page to go to
  if Language_user == "farsi" {

    if Language_to_learn == "german" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_farsi_learn_german.html"))
      tmpl.Execute(w, nil)
    } else if Language_to_learn == "french" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_farsi_learn_french.html"))
      tmpl.Execute(w, nil)
    }

  } else if Language_user == "german" {
    if Language_to_learn == "farsi" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_german_learn_farsi.html"))
      tmpl.Execute(w, nil)
    } else if Language_to_learn == "french" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_german_learn_french.html"))
      tmpl.Execute(w, nil)
    }

  } else if Language_user == "french" {
    if Language_to_learn == "farsi" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_french_learn_farsi.html"))
      tmpl.Execute(w, nil)
    } else if Language_to_learn == "german" {
      tmpl := template.Must(template.ParseFiles("html_files/" + Language_user+ "/startPage_user_french_learn_german.html"))
      tmpl.Execute(w, nil)
    }
  }

}



func startPageActionPost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  Language_url := Language_user + "To" + Language_to_learn

  switch buttonVal.Val {
  // Handle method
  case "wordList":
    http.Redirect(w, r, "/" + Language_url + "/listWords", http.StatusFound)
  case "wordTrainer":
    http.Redirect(w, r, "/" + Language_url + "/wordTrain", http.StatusFound)
  case "imageTrainer":
    http.Redirect(w, r, "/" + Language_url + "/imageTrain", http.StatusFound)
  }
}
