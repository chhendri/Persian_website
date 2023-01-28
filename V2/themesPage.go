package main

import (
  "net/http"
  "fmt"
  "strings"
  //"encoding/json"
)


type Theme struct {
  ID      int     `json:"id"`
  Title   string  `json:"title"`
  Wording string  `json:"wording"`
  Image   string  `json:"image"`
}


func chooseThemeGet (w http.ResponseWriter, r *http.Request) {
  inputs := parseURL(r.URL.Path)
  userLang := inputs[0]
  learnLang := inputs[1]
  levelLang := inputs[2]

  // Set the header and the text of the start page
  header := headerWording(userLang)
  title := defineTitle(userLang)
  themes, err := getAllThemes(userLang, learnLang, levelLang)
  if err != nil {
    fmt.Errorf("Error : Data of themes not fetched %v", err)
  }
  data := struct {
          Title    string
          Themes   []Theme
          Header   Header
  }{
          title,
          themes,
          header,
  }
  templates.ExecuteTemplate(w, "choose_theme.html", data)
}

func chooseThemePost (w http.ResponseWriter, r *http.Request) {
  // Parse the user input
  theme := r.FormValue("Theme")
  //Redirect the user depending on the input
  newUrl := r.URL.Path + "/" + string(theme)
  http.Redirect(w, r, newUrl, http.StatusFound)
}

func parseURL (url string) (inputs []string){
  // Parse the URL
  inputs = strings.Split(url[8:len(url)], "/")
  return inputs
}

func defineTitle (userLang string) (title string) {
  // Define the title according to the language of the user
  if userLang == "french" {
    title = "Choisissez un thème"
  } else if userLang == "english" {
    title = "Choose a theme"
  } else if userLang == "dutch" {
    title = "Kies een thema"
  } else if userLang == "german" {
    title = "Wähle ein thema"
  } else if userLang == "persian" {
    title = "Persian : Choose a theme"
  }

  return title
}

func getAllThemes (userLang string, learnLang string, levelLang string) ([]Theme, error) {
  var themes []Theme
  // Query to the database
  rows, err := db.Query("SELECT THEME_ID, TITLE_" + strings.ToUpper(userLang) + ", WORDING_" + strings.ToUpper(userLang) + ", IMAGE_ID FROM Themes")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Theme data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var theme Theme
    if err := rows.Scan(&theme.ID, &theme.Title, &theme.Wording, &theme.Image); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Theme data:  %v", err)
    }
    themes = append(themes, theme)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Theme data:   %v", err)
  }
  return themes, err
}
