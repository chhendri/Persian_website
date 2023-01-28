package main

import (
  "net/http"
  "strings"
  "fmt"
)


type userInput struct {
  userLang   string `json:"userLang"`
  learnLang  string `json:"learnLang"`
  levelLang  string `json:"levelLang"`
}

type Languages struct {
  French   string   `json:"french"`
  English  string   `json:"english"`
  Dutch    string   `json:"dutch"`
  German   string   `json:"german"`
  Persian  string   `json:"persian"`
}

type LangLevel struct {
  Level    string   `json:"level"`
  Wording  string   `json:"wording"`
}


func renderChooseParamsPage(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path)
  // Choose the user language
  userLang := strings.TrimPrefix(r.URL.Path, "/lang/")
  userLang = strings.TrimSuffix(userLang, "/chooseParams")

  // Get the text for the Params Page
  header := headerWording(userLang)
  title, subtitle_1, subtitle_2, subtitle_3, button := wordingParams(userLang)
  languages := wordingLanguages(userLang)
  levels, err := getLanguageLevels(userLang)
  if err != nil {
    fmt.Errorf("Error : Data of levels not fetched %v", err)
  }

  // Create the data structures
  data := struct {
    Header      Header
    Title       string
    Subtitle_1  string
    Subtitle_2  string
    Subtitle_3  string
    Languages   Languages
    Button      string
    LangLevels  []LangLevel
  }{
    header,
    title,
    subtitle_1,
    subtitle_2,
    subtitle_3,
    languages,
    button,
    levels,
  }
  templates.ExecuteTemplate(w, "choose_params.html", data)
}

func userParams (w http.ResponseWriter, r *http.Request) {
  // Parse the user input
  var userInp userInput
  userInp.userLang = r.FormValue("userLang")
  userInp.learnLang = r.FormValue("learnLang")
  userInp.levelLang = r.FormValue("levelLang")
  //Redirect the user depending on the input
  newUrl := "/themes/" + userInp.userLang + "/" + userInp.learnLang + "/" + userInp.levelLang
  http.Redirect(w, r, newUrl, http.StatusFound)
}

func wordingLanguages (userLang string) (langWording Languages) {
  if userLang == "french" {
    langWording = Languages {
       "Français",
       "Anglais",
       "Néerlandais",
       "Allemand",
       "Persan",
     }
  } else if userLang == "english" {
    langWording = Languages {
       "French",
       "English",
       "Dutch",
       "German",
       "Persian",
     }
  } else if userLang == "dutch" {
    langWording = Languages {
       "Frans",
       "Engels",
       "Nederlands",
       "Duits",
       "Persisch",
     }
  } else if userLang == "german" {
    langWording = Languages {
       "Französisch",
       "Englisch",
       "Niederländisch",
       "Deutsch",
       "Persisch",
     }
  } else if userLang == "persian" {
    langWording = Languages {
       "French",
       "English",
       "Dutch",
       "German",
       "Persian",
     }
  }
  return langWording
}

func wordingParams (userLang string) (title string, subtitle_1 string, subtitle_2 string, subtitle_3 string, button string) {
  if userLang == "french" {
    title = "Choisissez"
    subtitle_1 = "Une langue que vous conaissez bien"
    subtitle_2 = "La langue que vous voulez apprendre"
    subtitle_3 = "Votre niveau dans la langue que vous voulez apprendre"
    button = "Commencez"
  } else if userLang == "english" {
    title = "Choose"
    subtitle_1 = "A language you know well"
    subtitle_2 = "The language you want to learn"
    subtitle_3 = "Your level in the language you want to learn"
    button = "Get started"
  } else if userLang == "dutch" {
    title = "Kies"
    subtitle_1 = "Een taal dat u goed kent"
    subtitle_2 = "De taal dat u leren wil"
    subtitle_3 = "Uw niveau in de taal die u leren wil"
    button = "Aan de slag"
  } else if userLang == "german" {
    title = "Wähle"
    subtitle_1 = "Ein Sprache dem sie gut kennen"
    subtitle_2 = "Der Sprache dem sie lernen wollen"
    subtitle_3 = "Ihrem Niveau in die Sprache dem sie lernen wollen"
    button = "Los geht's"
  } else if userLang == "persian" {
    title = "Choose"
    subtitle_1 = "A language you know well"
    subtitle_2 = "The language you want to learn"
    subtitle_3 = "Your level in the language you want to learn"
    button = "Get started"
  }
  return title, subtitle_1, subtitle_2, subtitle_3, button
}

func getLanguageLevels (userLang string) ([]LangLevel, error) {
  var langLevels []LangLevel

  // Query to the database
  rows, err := db.Query("SELECT LEVEL, WORDING_" + strings.ToUpper(userLang) + " FROM Levels where LEVEL in (select distinct LEVEL from Themes)")
  if err != nil {
    return nil, fmt.Errorf("Error : Cannot fetch Levels data: %v", err)
  }

  defer rows.Close()
  // Store the results in the Themes object
  for rows.Next() {
    var level LangLevel
    if err := rows.Scan(&level.Level, &level.Wording); err != nil {
      return nil, fmt.Errorf("Error : Cannot scan Levels data:  %v", err)
    }
    langLevels = append(langLevels, level)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("Error : Not able to format Levels data:   %v", err)
  }

  return langLevels, nil
}
