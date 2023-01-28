package main

import (
  "fmt"
  "net/http"
  "strings"
  "math/rand"
  "time"
  "strconv"
)


type Word_Correct struct {
  ID      string   `json:"id"`
  wULang  string   `json:"Word_User_Lang"`
  wLLang  string   `json:"Word_Learn_Lang"`
}

type Word_Corr_Compare struct {
  ID              string   `json:"id"` // The ID of the Word in the database
  IsCorrect       bool     `json:"correct"` // If the input of the user is correct
  UserInput       string   `json:"user_input"` // Input of the user
  WUserCorrect    string   `json:"word_user"` // Word that should have been inputted
  WLearnCorrect   string   `json:"word_learn"` // Translation of the word
}


func getWords (userLang string, learnLang string, themeNbr string) ([]Word, error) {
  var words []Word
  // Query to the database
  rows, err := db.Query("SELECT " + strings.ToUpper(learnLang) + ", " + strings.ToUpper(userLang) + ", WORD_ID, IMAGE_ID FROM Words WHERE THEME_ID='" + themeNbr + "'")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Word data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var word Word
    if err := rows.Scan(&word.WordULang, &word.WordLLang, &word.ID, &word.Image); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Word data:  %v", err)
    }
    words = append(words, word)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Word data:   %v", err)
  }
  return words, err
}

func getWordsImages(userLang string, learnLang string, themeNbr string) ([]Word, error) {
  var words []Word
  // Query to the database
  rows, err := db.Query("SELECT " + strings.ToUpper(learnLang) + ", " + strings.ToUpper(userLang) + ", WORD_ID, IMAGE_ID FROM Words WHERE THEME_ID='" + themeNbr + "' AND IMAGE_ID != ''")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Word data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var word Word
    if err := rows.Scan(&word.WordULang, &word.WordLLang, &word.ID, &word.Image); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Word data:  %v", err)
    }
    words = append(words, word)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Word data:   %v", err)
  }
  return words, err
}

func getDialogue (userLang string, learnLang string, themeNbr string) ([]Dialogue, error) {
  var dialogues []Dialogue
  // Query to the database
  rows, err := db.Query("SELECT DIALOGUE_ID, SEQUENCE_NUMBER, CHARACTER_" + strings.ToUpper(learnLang) + ", " + strings.ToUpper(learnLang) + ", CHARACTER_" + strings.ToUpper(userLang) + ", " + strings.ToUpper(userLang) + " FROM Dialogues WHERE THEME_ID='" + themeNbr + "' ORDER BY DIALOGUE_ID, SEQUENCE_NUMBER ASC")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Dialogue data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var dialogue Dialogue
    if err := rows.Scan(&dialogue.DiagID, &dialogue.SeqNbr, &dialogue.CharLLang, &dialogue.DiagLLang, &dialogue.CharULang, &dialogue.DiagULang); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Dialogue data:  %v", err)
    }
    dialogues = append(dialogues, dialogue)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Dialogue data:   %v", err)
  }
  return dialogues, err
}

func tabChangePost (w http.ResponseWriter, r *http.Request) {
  var action string
  // Parse the user input
  r.ParseForm()
  for key := range r.PostForm {
    action = key
  }
  //Redirect the user depending on the input
  inputs := parseURL(r.URL.Path)[0:4]
  newUrl := "/themes/" + strings.Join(inputs, "/") + "/" + string(action)
  http.Redirect(w, r, newUrl, http.StatusFound)
}

func distribWordsExercices (words_in []Word) (words_ex1 []Word, words_ex2 []Word, anonymWords []Word_Ex3, correctWords []Word_Ex3) {
  var words_out []Word
  // Distribute the words depending on the exercice
  rand.Seed(time.Now().UnixNano())
  min := 1
  max := len(words_in)
  // Select random words
  for i:=0; i < 35; i++ {
    r := rand.Intn(max-min) + min
    words_out = append(words_out, words_in[r])
  }
  // Divide the words between the exercices
  words_ex1 = words_out[0:10]
  words_ex2 = words_out[10:20]
  anonymWords, correctWords = shuffleWordsEx3(words_out)
  return words_ex1, words_ex2, anonymWords, correctWords
}

func shuffleWordsEx3 (words_out []Word) (anonymWords []Word_Ex3, correctWords []Word_Ex3) {
  var LLang_IDs []shuffleID
  // For Exercice 3, we need to define a new ID in order to match the words
  letter_ids := randomLetters(15)
  for i:=20; i < 35; i++ {
    var correctWord Word_Ex3
    correctWord.WordULang = words_out[i].WordULang
    correctWord.WordLLang = words_out[i].WordLLang
    correctWord.ID_ULang = strconv.Itoa(i-20)
    correctWord.ID_LLang = letter_ids[i-20]
    correctWord.ID = words_out[i].ID
    correctWords = append(correctWords, correctWord)
    // Get the ID_LLang of the elements
    var new shuffleID
    new.WordLLang = correctWord.WordLLang
    new.ID = correctWord.ID_LLang
    LLang_IDs = append(LLang_IDs, new)
  }
  // Shuffle ID_LLang
  for i:=0; i<len(LLang_IDs); i++ {
    j := rand.Intn(i + 1)
    LLang_IDs[i], LLang_IDs[j] = LLang_IDs[j], LLang_IDs[i]
  }
  // Assign the new ID_LLang
  for i:=0; i<len(correctWords); i++ {
    anonymWord := correctWords[i]
    anonymWord.WordLLang = LLang_IDs[i].WordLLang
    anonymWord.ID_LLang = LLang_IDs[i].ID
    anonymWords = append(anonymWords, anonymWord)
  }
  return anonymWords, correctWords
}

func randomLetters (max int) (letters []string) {
  poss_letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  letters = strings.SplitAfter(poss_letters[0:max], "")
  return letters
}

func getWordsByID (IDs []string, userLang string, learnLang string) ([]Word_Correct, error) {
  var Words_Correct []Word_Correct
  // Query to the database
  rows, err := db.Query("SELECT WORD_ID, " + strings.ToUpper(userLang) + ", " + strings.ToUpper(learnLang) + " FROM Words WHERE WORD_ID in (" + strings.Join(IDs, ", ") + ")")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Correction data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var wCorr Word_Correct
    if err := rows.Scan(&wCorr.ID, &wCorr.wULang, &wCorr.wLLang); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Correction data:  %v", err)
    }
    Words_Correct = append(Words_Correct, wCorr)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Correction data:   %v", err)
  }

  return Words_Correct, err
}

func removeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func compareInputCorrect (IDs []string, userInput []Word_Correct, wordsCorr []Word_Correct) (compEX1 []Word_Corr_Compare, compEX2 []Word_Corr_Compare, compEX3 []Word_Corr_Compare) {
  // Loop over the user input
  for i := range userInput {
    // Number of the exercice
    Ex := userInput[i].ID[0:3]
    // ID of the word asked
    ID := strings.TrimPrefix(userInput[i].ID, Ex + "_")
    // Word inputted by the user
    uImp := userInput[i].wULang
    // Loop over the correct words
    for i := range wordsCorr {
      ID_corr := wordsCorr[i].ID
      // Find the corresponding word
      if ID == ID_corr {
        var correct bool
        // Check if it is correct. We remove white spaces to make it clearer.
        if removeSpaces(uImp) == removeSpaces(wordsCorr[i].wULang) {
          correct = true
        } else if removeSpaces(uImp) == removeSpaces(wordsCorr[i].wLLang) {
          correct = true
        } else {
          correct = false
        }
        // Set the output
        word := Word_Corr_Compare {
          ID,
          correct,
          uImp,
          wordsCorr[i].wULang,
          wordsCorr[i].wLLang,
        }
        // Add the output to the correct exercice
        if Ex == "ex1" {
          compEX1 = append(compEX1, word)
        } else if Ex == "ex2" {
          compEX2 = append(compEX2, word)
        } else if Ex == "ex3" {
          compEX3 = append(compEX3, word)
        }
      }
    }
  }
  return compEX1, compEX2, compEX3
}
