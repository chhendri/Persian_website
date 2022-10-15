package main

import (
  "net/http"
  "time"
  "math/rand"
  "html/template"
  "strings"
  "strconv"
)


type Woid struct {
  ID int `json:"id"`
}



type correctionWords struct {
  QueryWord    string   `json:"queryWord"`
  UserWord     string   `json:"userWord"`
  CorrectWord  string   `json:"correctWord"`
  Color        string   `json:"color"`
}



type inputWords struct {
  Title string `json:"Title"`
  SubTitle string `json:"SubTitle"`
  Words []correctionWords  `json:"Words"`
}


var queryWord string
var correctTrans string
var randWoid Woid
var woidsToDo []Woid
var userWord string


// Functions for display of Lecture selector
func wordTrainerGet(user_language string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {

    connectDB(w, r)
    rand.Seed(time.Now().UnixNano())

    var Titles inputWords

    // Set the title of the page
    if user_language == "farsi" {
      Titles.Title = "Choose your lecture"
      Titles.SubTitle = "Lecture"
    } else if user_language == "french" {
      Titles.Title = "Choisissez votre cours"
      Titles.SubTitle = "Cours"
    } else if user_language == "german" {
      Titles.Title = "Wahle ihrem Kurs"
      Titles.SubTitle = "Kurs"
    }

    // Display the html
    tmpl := template.Must(template.ParseFiles("html_files/wordTrainer.html"))
    tmpl.Execute(w, Titles)
  }
}


func wordTrainerPost(w http.ResponseWriter, r *http.Request)  {
  var buttonVal buttonValue
  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }
  http.Redirect(w, r, r.URL.Path + "/" + buttonVal.Val, http.StatusFound)
}



// Function for query of all words of lecture
func wordTrainLectureGet(leid string, user_language string, to_learn_language string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

      var corrList inputWords

      // Set the title of the page
      if user_language == "farsi" {
        corrList.Title = "Train your vocabulary"
        corrList.SubTitle = "For each word, enter the correct translation and click on submit to get the correction."
      } else if user_language == "french" {
        corrList.Title = "Entrainez votre vocabulaire"
        corrList.SubTitle = "Pour chaque mot, entrez la traduction correcte."
      } else if user_language == "german" {
        corrList.Title = "Übe ihrem Vokabular"
        corrList.SubTitle = "Für jedem Wort, setze dem korrekte Überzetsung"
      }

      tmpl_lec := template.Must(template.ParseFiles("html_files/" + user_language+ "/wordTrainerLecture_" + leid + "_" + user_language + "_" + to_learn_language + ".html"))
      tmpl_lec.Execute(w, corrList)
    }
}


func wordTrainLecturePost(leid string, user_language string, to_learn_language string) http.HandlerFunc {
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {
        //words := wordsByLecture(leid, user_language, to_learn_language)
        r.ParseForm()

        // Get words of lecture
        //words := wordsByLecture(leid, user_language, to_learn_language)

        var corrList inputWords

        // Set the title of the page
        if user_language == "farsi" {
          corrList.Title = "Train your vocabulary"
          corrList.SubTitle = "The correction of the exercice"
        } else if user_language == "french" {
          corrList.Title = "Entrainez votre vocabulaire"
          corrList.SubTitle = "La correction de l'exercice"
        } else if user_language == "german" {
          corrList.Title = "Übe ihrem Vokabular"
          corrList.SubTitle = "Dem Korrektur von dem Übung"
        }
        /*
        for _ , i := range words {
          var wordCorr correctionWords{
            {queryWord: i.Fran, userWord:, correctWord: i.Pers}
          }
        }
        */
        for key := range r.PostForm {
          if key != "submit" {
            // Get the Woid
            woid, _ := strconv.Atoi(strings.TrimPrefix(key, "response_"))
            corrOne := wordsByWoid(woid, user_language, to_learn_language)
            corrOne.UserWord = r.FormValue(key)
            if corrOne.UserWord != corrOne.CorrectWord {
              corrOne.Color = "#ba0f30"
            } else if corrOne.UserWord == corrOne.CorrectWord {
              corrOne.Color = "#2a7221"
            }
            corrList.Words = append(corrList.Words, corrOne)
          }
        }

        htmlName := "wordTrainerCorrLecture_" + leid + "_" + user_language + "_" + to_learn_language + ".html"
        tmpl_lec := template.Must(template.ParseFiles("html_files/" + user_language+ "/" + htmlName))
        tmpl_lec.Execute(w, corrList)
    }
}


func findWordInArray(woid int, words []Word) (word Word) {
  // Find the word if it is in an array
  for _, w := range words {
    if w.Woid == woid {
      return w
    }
  }
  var emptyWord Word
  return emptyWord
}
