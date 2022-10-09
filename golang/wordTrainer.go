package main

import (
  "fmt"
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
  Woid         int      `json:"Woid"`
  QueryWord    string   `json:"queryWord"`
  UserWord     string   `json:"userWord"`
  CorrectWord  string   `json:"correctWord"`
}

var queryWord string
var correctTrans string
var randWoid Woid
var woidsToDo []Woid
var userWord string


// Functions for display of Lecture selector

func wordTrainerGet(w http.ResponseWriter, r *http.Request)  {
  connectDB(w, r)
  rand.Seed(time.Now().UnixNano())
  // Display the html
  tmpl := template.Must(template.ParseFiles("html_files/wordTrainer.html"))
  tmpl.Execute(w, nil)
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
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {
      tmpl_lec := template.Must(template.ParseFiles("html_files/" + user_language+ "/wordTrainerLecture_" + leid + "_" + user_language + "_" + to_learn_language + ".html"))
      tmpl_lec.Execute(w, nil)
    }
}


func wordTrainLecturePost(leid string, user_language string, to_learn_language string) http.HandlerFunc {
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {
        words := wordsByLecture(leid, user_language, to_learn_language)
        r.ParseForm()
        var corrList []correctionWords
        for key := range r.PostForm {
          if key != "submit" {
            var corrWor correctionWords
            // User input
            corrWor.UserWord = r.FormValue(key)
            // Get wordID from button name
            corrWor.Woid, _ = strconv.Atoi(strings.TrimPrefix(key, "response_"))
            // Get correct word
            wor := findWordInArray(corrWor.Woid, words)
            corrWor.QueryWord = wor.Fran
            corrWor.CorrectWord = wor.Pers
            corrList = append(corrList, corrWor)
            fmt.Println(corrWor)
          }
        }
        htmlName := constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, user_language, to_learn_language, corrList)
        tmpl_lec := template.Must(template.ParseFiles("html_files/" + user_language+ "/" + htmlName))
        tmpl_lec.Execute(w, nil)
    }
}





/*
func askWordVisualGet(w http.ResponseWriter, r *http.Request){
  // Get the words of that lecture
  words := wordsByLecture("1", Language_user, Language_to_learn)

  // Get the woid of the words
  woids := woidsFromWords(words)
  // List of words already queried
  woidsToDo := woids
  queryWord, correctTrans, randWoid := askWord(words, woidsToDo)
  fmt.Println("correcttrans= " + correctTrans)
  fmt.Println("queryword= " + queryWord)
  woidsToDo = removeWordFromToDo(randWoid, woidsToDo)
  tmpl := template.Must(template.ParseFiles("html_files/trainLecture1.html"))
  tmpl.Execute(w, queryWord)
}


func askWordVisualPost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue
  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }
  userWord := r.FormValue("trans")
  // Here problem because the correctTrans is not available to this function
  fmt.Println("userword = " + userWord)
  fmt.Println("correctTrans2= " + correctTrans)
  if correctTrans == userWord {
    http.Redirect(w, r, "/wordTrain/correctWord", http.StatusFound)
  } else {
    http.Redirect(w, r, "/wordTrain/falseWord", http.StatusFound)
  }
}



func htmlForFalsePost(queryWord string, userWord string, correctTrans string) http.HandlerFunc {
  var corrWord correctionWords
  corrWord.QueryWord = queryWord
  corrWord.UserWord = userWord
  corrWord.CorrectWord = correctTrans
  return func(w http.ResponseWriter, r *http.Request) {
    tmpl_lec := template.Must(template.ParseFiles("html_files/error_wordTrain.html"))
    tmpl_lec.Execute(w, corrWord)
  }
}


func htmlForCorrectPost(queryWord string, userWord string, correctTrans string) http.HandlerFunc {
  // Display the html
  var corrWord correctionWords
  corrWord.QueryWord = queryWord
  corrWord.UserWord = userWord
  corrWord.CorrectWord = correctTrans
  return func(w http.ResponseWriter, r *http.Request) {
    tmpl_lec := template.Must(template.ParseFiles("html_files/correct_wordTrain.html"))
    tmpl_lec.Execute(w, corrWord)
  }
}



// Functions for querying the word

func woidsFromWords(words []Word) (woids []Woid) {
  // Get the Woid from all words
  for _ , i := range words {
    var woid Woid
    woid.ID = i.Woid
    woids = append(woids, woid)
  }
  return woids
}


func selectRandomWoid (woids []Woid) (randWoid Woid){
  // Take a random woid from the list of woids
  randomIndex := rand.Intn(len(woids))
  randWoid = woids[randomIndex]
  return randWoid
}
*/


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

/*
func removeWordFromToDo (woid Woid, woidsToDo []Woid) (woidsToDo2 []Woid) {
  // Remove word from a list
  for i := len(woidsToDo) - 1; i >= 0; i-- {
    w := woidsToDo[i]
    // Condition to decide if current element has to be deleted:
    if w == woid {
        woidsToDo2 = append(woidsToDo[:i],
                woidsToDo[i+1:]...)
    }
  }
  return woidsToDo2
}


func askWord (words []Word, woidsToDo []Woid) (qword string, rword string, randWoid Woid) {
  // Select a random word
  randWoid = selectRandomWoid(woidsToDo)
  word := findWordInArray(randWoid, words)
  return word.Fran, word.Pers, randWoid
}
*/
