package main

import(
  "fmt"
  "net/http"
  "strings"
)


type Word struct {
    WordULang   string `json:"wordULang"`
    WordLLang   string `json:"worLLang"`
    Image       string `json:"image"`
    ID          string `json:"id"`
}


type Word_Ex3 struct {
    WordULang   string `json:"wordULang"`
    WordLLang   string `json:"worLLang"`
    ID_ULang    string `json:"id_ULang"`
    ID_LLang    string `json:"id_LLang"`
    ID          string `json:"id"`
}


type Dialogue struct {
  DiagID       string `json:"diagID"`
  SeqNbr       string `json:"seqNbr"`
  CharULang    string `json:"charULang"`
  CharLLang    string `json:"charLLang"`
  DiagULang    string `json:"diagULang"`
  DiagLLang    string `json:"diagLLang"`
}

type shuffleID struct {
  ID        string
  WordLLang string
}

func trainVocabularyGet (w http.ResponseWriter, r *http.Request) {
  // Parse the URL
  inputs := parseURL(r.URL.Path)
  userLang := inputs[0]
  learnLang := inputs[1]
  themeNbr := inputs[3]
  // Set the header and the text of the start page
  header := headerWording(userLang)
  // Find the wording for the HTML page
  voc_title, diag_title, ex_image_title, ex_word_title := defineTabWording(userLang)
  // Get the words of the theme
  words, err := getWords(userLang, learnLang, themeNbr)
  if err != nil {
    fmt.Errorf("Error : Data of words not fetched %v", err)
  }
  // Create the data
  data := struct {
          Header          Header
          Vocabulary      string
          Dialogue        string
          Exercice_Image  string
          Exercice_Words  string
          Words           []Word
  }{
          header,
          voc_title,
          diag_title,
          ex_image_title,
          ex_word_title,
          words,
  }
  // Render the HTML page
  templates.ExecuteTemplate(w, "train_page_vocabulary.html", data)
}

func trainDialogueGet (w http.ResponseWriter, r *http.Request) {
  // Parse the URL
  inputs := parseURL(r.URL.Path)
  userLang := inputs[0]
  learnLang := inputs[1]
  themeNbr := inputs[3]
  // Set the header and the text of the start page
  header := headerWording(userLang)
  // Find the wording for tgetWordshe HTML page
  voc_title, diag_title, ex_image_title, ex_word_title := defineTabWording(userLang)
  // Get the dialogue of the theme
  dialogues, err := getDialogue(userLang, learnLang, themeNbr)
  if err != nil {
    fmt.Errorf("Error : Data of words not fetched %v", err)
  }
  // Create the data
  data := struct {
          Header          Header
          Vocabulary      string
          Dialogue        string
          Exercice_Image  string
          Exercice_Words  string
          Dialogues       []Dialogue
  }{
          header,
          voc_title,
          diag_title,
          ex_image_title,
          ex_word_title,
          dialogues,
  }
  // Render the HTML page
  templates.ExecuteTemplate(w, "train_page_dialogue.html", data)
}

func trainWordsGet (w http.ResponseWriter, r *http.Request) {
  // Parse the URL
  inputs := parseURL(r.URL.Path)
  userLang := inputs[0]
  learnLang := inputs[1]
  themeNbr := inputs[3]
  // Set the header and the text of the start page
  header := headerWording(userLang)
  // Find the wording for the HTML page
  voc_title, diag_title, ex_image_title, ex_word_title := defineTabWording(userLang)
  corr_title, ex1_ex2_title, ex3_title := defineExerciceWording(userLang)
  // Get the words of the theme
  words, err := getWords(userLang, learnLang, themeNbr)
  if err != nil {
    fmt.Errorf("Error : Data of words not fetched %v", err)
  }
  words_ex1, words_ex2, anonymWords, correctWords := distribWordsExercices(words)
  _ = correctWords
  // Create the data
  data := struct {
          Header          Header
          Vocabulary      string
          Dialogue        string
          Exercice_Image  string
          Exercice_Words  string
          Correct         string
          EX1_EX2_Title   string
          EX3_Title       string
          Words_Ex1       []Word
          Words_Ex2       []Word
          Words_Ex3       []Word_Ex3
  }{
          header,
          voc_title,
          diag_title,
          ex_image_title,
          ex_word_title,
          corr_title,
          ex1_ex2_title,
          ex3_title,
          words_ex1,
          words_ex2,
          anonymWords,
  }
  // Render the HTML page
  templates.ExecuteTemplate(w, "train_page_words.html", data)
}


func trainImagesGet (w http.ResponseWriter, r *http.Request) {
  // Parse the URL
  inputs := parseURL(r.URL.Path)
  userLang := inputs[0]
  learnLang := inputs[1]
  themeNbr := inputs[3]
  // Find the wording for the HTML page
  voc_title, diag_title, ex_image_title, ex_word_title := defineTabWording(userLang)
  // Set the header and the text of the start page
  header := headerWording(userLang)
  // Get the words of the theme
  words, err := getWordsImages(userLang, learnLang, themeNbr)
  if err != nil {
    fmt.Errorf("Error : Data of words of images not fetched %v", err)
  }
  // Create the data
  data := struct {
          Header          Header
          Vocabulary      string
          Dialogue        string
          Exercice_Image  string
          Exercice_Words  string
          Words           []Word
  }{
          header,
          voc_title,
          diag_title,
          ex_image_title,
          ex_word_title,
          words,
  }

  templates.ExecuteTemplate(w, "train_page_images.html", data)
}


func trainWordsCorrect (w http.ResponseWriter, r *http.Request) {
  // Get the user and learn language
  url_split := strings.Split(r.URL.Path, "/")
  userLang := url_split[2]
  learnLang := url_split[3]

  // Find the user input
  var IDs        []string
  var userInputs []Word_Correct

  values := r.URL.Query()
  for k, v := range values {
    if k != "CorrectWords" {
      IDs = append(IDs, strings.TrimPrefix(k, k[0:3] + "_"))
      userInputs = append(userInputs, Word_Correct{k, strings.Join(v," "), ""})
    }
  }
  // Get the correct words
  words_corr, err := getWordsByID(IDs, userLang, learnLang)
  if err != nil {
    fmt.Errorf("Error : Cannot fetch Correction data: %v", err)
  }
  // Compare the user input with the correct words
  compEX1, compEX2, compEX3 := compareInputCorrect(IDs, userInputs, words_corr)
  // Set the header and the text of the start page
  header := headerWording(userLang)
  // Find the wording for the HTML page
  voc_title, diag_title, ex_image_title, ex_word_title := defineTabWording(userLang)
  corr_title, ex1_ex2_title, ex3_title := defineExerciceWording(userLang)
  _ = corr_title

  // Create the data
  data := struct {
          Header          Header
          Vocabulary      string
          Dialogue        string
          Exercice_Image  string
          Exercice_Words  string
          EX1_EX2_Title   string
          EX3_Title       string
          WordsEx1        []Word_Corr_Compare
          WordsEx2        []Word_Corr_Compare
          WordsEx3        []Word_Corr_Compare
  }{
          header,
          voc_title,
          diag_title,
          ex_image_title,
          ex_word_title,
          ex1_ex2_title,
          ex3_title,
          compEX1,
          compEX2,
          compEX3,
  }
  templates.ExecuteTemplate(w, "train_page_words_correct.html", data)
}
