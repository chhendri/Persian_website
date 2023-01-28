package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

var Serve http.Handler


func Mux() {
    r := mux.NewRouter()

    r.HandleFunc("/", redirectStartPage)
    r.HandleFunc("/lang/{uLang}", renderStartPage)
    r.HandleFunc("/lang/{uLang}/chooseParams", renderChooseParamsPage).Methods("GET")
    r.HandleFunc("/lang/{uLang}/chooseParams", userParams).Methods("POST")
    r.HandleFunc("/lang/{uLang}/contact", renderContactPage)
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}", chooseThemeGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}", chooseThemePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}", trainVocabularyGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}", tabChangePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Vocabulary", trainVocabularyGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Vocabulary", tabChangePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Dialogue", trainDialogueGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Dialogue", tabChangePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Exercice_Words", trainWordsGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Exercice_Words", tabChangePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/CorrectWords", trainWordsCorrect).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/CorrectWords", tabChangePost).Methods("POST")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Exercice_Image", trainImagesGet).Methods("GET")
    r.HandleFunc("/themes/{uLang}/{lLang}/{vLang}/{themeID}/Exercice_Image", tabChangePost).Methods("POST")

    Serve = r
    http.ListenAndServe(":8000", Serve)
}
