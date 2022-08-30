package main

import (
  "fmt"
  "net/http"
  "time"
  "math/rand"
  "html/template"
)


type Woid struct {
  ID int `json:"id"`
}


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
  http.Redirect(w, r, "/wordTrain/" + buttonVal.Val, http.StatusFound)
}


func askWordVisual(w http.ResponseWriter, r *http.Request){
  // Get the words of that lecture
  words, err := wordsByLecture("1")
  if err != nil {
    fmt.Println("There was an error")
  }
  // Get the woid of the words
  woids := woidsFromWords(words)
  // List of words already queried
  woidsToDo := woids
  i, j, randWoid := askWord(words, woidsToDo)
  htmlForWord(w, r, i)
  woidsToDo = removeWordFromToDo(randWoid, woidsToDo)
  // Check input of user
  fmt.Println(j)
}


func htmlForWord(w http.ResponseWriter, r *http.Request, word string){
  // Display the html
  tmpl := template.Must(template.ParseFiles("html_files/trainLecture1.html"))
  tmpl.Execute(w, word)
}


func submitCorrect(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }
  fmt.Println("buttonval: " + buttonVal.Val)
  http.Redirect(w, r, "/listWords", http.StatusFound)
}


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


func findWordInArray(woid Woid, words []Word) (word Word) {
  // Find the word if it is in an array
  for _, w := range words {
    if w.Woid == woid.ID {
      return w
    }
  }
  var emptyWord Word
  return emptyWord
}


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
