package main

import (
  "net/http"
  "strings"
  "fmt"
)

type Button struct {
  ID      int     `json:"id"`
  Value   string  `json:"value"`
}

type Header struct {
  Home        string     `json:"home"`
  Language    string     `json:"language"`
  Navigation  string     `json:"navigation"`
  Contact     string     `json:"contact"`
  French   string   `json:"french"`
  English  string   `json:"english"`
  Dutch    string   `json:"dutch"`
  German   string   `json:"german"`
  Persian  string   `json:"persian"`
}


func redirectStartPage(w http.ResponseWriter, r *http.Request){
  http.Redirect(w, r, "/lang/english", http.StatusSeeOther)
}


func renderStartPage(w http.ResponseWriter, r *http.Request) {
  // Choose the user language
  var userLang string
  if r.URL.Path == "/" {
    userLang = "english"
  } else {
    userLang = strings.TrimPrefix(r.URL.Path, "/lang/")
  }
  // Set the header and the text of the start page
  header := headerWording(userLang)
  title, subtitle, button := startPageWording(userLang)
  // Construct the data structure to pass through
  data := struct {
      Header    Header
      Title     string
      Subtitle  string
      Button    string
    }{
      header,
      title,
      subtitle,
      button,
    }
  templates.ExecuteTemplate(w, "start_page.html", data)
}

func redirectLanguage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Change language")
}

func startPageWording(userLang string) (title string, subtitle string, button string) {
  if userLang == "french" {
    title = "Apprenez du vocabulaire de cinq langues"
    subtitle = "Gratuitement"
    button = "Commencez"
  } else if userLang == "english" || userLang == "" {
    title = "Learn vocabulary of five languages"
    subtitle = "for Free"
    button = "Get started"
  } else if userLang == "dutch" {
    title = "Leer woordenschat van vijf talen"
    subtitle = "Gratis"
    button = "aan de Slag"
  } else if userLang == "german" {
    title = "Lerne Wortschatz von funf Sprachen"
    subtitle = "Kostenlos"
    button = "Los geht's"
  } else if userLang == "persian" {
    title = "Learn vocabulary of five languages"
    subtitle = "for Free"
    button = "Get started"
  }
  return title, subtitle, button
}


func headerWording(userLang string) (header Header) {
  if userLang == "french" {
    header = Header{
      "Acceuil",
      "Langue",
      "Navigation",
      "Contact",
      "Français",
      "Anglais",
      "Néerlandais",
      "Allemand",
      "Persan",
    }
  } else if userLang == "english" || userLang == "" {
    header = Header{
      "Home",
      "Language",
      "Navigation",
      "Contact",
      "French",
      "English",
      "Dutch",
      "German",
      "Persian",
    }
  } else if userLang == "dutch" {
    header = Header{
      "Home",
      "Taal",
      "Navigatie",
      "Contact",
      "Frans",
      "Engels",
      "Nederlands",
      "Duits",
      "Persisch",
    }
  } else if userLang == "german" {
    header = Header{
      "Home",
      "Sprache",
      "Navigation",
      "Kontakt",
      "Französisch",
      "Englisch",
      "Niederländisch",
      "Deutsch",
      "Persisch",
    }
  } else if userLang == "persian" {
    header = Header{
      "Home",
      "Language",
      "Navigation",
      "Contact",
      "French",
      "English",
      "Dutch",
      "German",
      "Persian",
    }
  }
  return header
}
