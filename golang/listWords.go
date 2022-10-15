package main

import (
  "net/http"
  "fmt"
  "log"
  //"github.com/gin-gonic/gin"
  "io/ioutil"
  "html/template"
)

type Word struct {
    Woid int    `json:"woid"`
    Fran string `json:"fran"`
    Tran string `json:"tran"`
    Pers string `json:"pers"`
    Leid int64  `json:"leid"`
    Imag string `json:"imag"`
}

type Lecture struct {
  ID int `json:"id"`
}


func listWordsGet(user_language string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {

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

  // Main function to display the listWords page
  tmpl := template.Must(template.ParseFiles("html_files/listWords.html"))
  tmpl.Execute(w, Titles)
  }
}


func listWordsPost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue
  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }
  http.Redirect(w, r, r.URL.Path + "/" + buttonVal.Val, http.StatusFound)
}


func lectureHandler(leid string, user_language string, to_learn_language string) http.HandlerFunc {
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {

      var Titles inputWords

      // Set the title of the page
      if user_language == "farsi" {
        Titles.Title = "Vocabular of lecture"
      } else if user_language == "french" {
        Titles.Title = "Vocabulaire du cours"
      } else if user_language == "german" {
        Titles.Title = "Vokabular von dem Kurs"
      }

      tmpl_lec := template.Must(template.ParseFiles("html_files/" + user_language+ "/Lecture" + leid + "_" + user_language + "_" + to_learn_language + ".html"))
      tmpl_lec.Execute(w, Titles)
    }
}


func tableVocabulary(jsonObj []Word, lecture_id string, user_language string, to_learn_language string) (fileName string){
  // Turn a JSON into an HTML file with header
  var htmlString string

  htmlString += "<!Doctype html>\n<html>\n<head>\n<meta charset='utf-8'>\n<title>listWords</title>\n</head>\n<body style='background-color:#e7ecef;'>\n"
  // Set the header for the Lecture file
  header := "<h1>{{ .Title }} " + lecture_id + "</h1><table> \n"
  // Set the header
  htmlString += header

  // Add the rows of the table corresponding to an instance of the JSON
  for _ , i := range jsonObj {
    htmlString += "<tr> \n"
    htmlString += "<td>" + string(i.Fran) + "</td> \n"
    htmlString += "<td>" + string(i.Pers) + "</td> \n"
    htmlString += "</tr> \n"
  }
  // Add the footer
  htmlString += "</table>\n</body>\n</html>"
  // Define the filename
  fileName += "Lecture" + lecture_id + "_" + user_language + "_" + to_learn_language + ".html"
  // Save to file
  if err := ioutil.WriteFile("html_files/" + user_language+ "/" + fileName, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }

  return fileName
}


func wordsByLecture(lecture string, user_language string, to_learn_language string) ([]Word) {
    // Get all the words for a lecture
    var words []Word
    // String giving the Query
    var query string

    // Query the database according to the scenario
    query = buildQueryWordList(user_language, to_learn_language)
    rows, err := db.Query(query, lecture)
    if err != nil {
        return nil
    }

    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var wor Word
        if err := rows.Scan(&wor.Woid, &wor.Fran, &wor.Pers, &wor.Leid, &wor.Imag); err != nil {
            return nil
        }
        words = append(words, wor)
    }

    return words
}


func buildQueryWordList(user_language string, to_learn_language string) (query string) {
  if user_language == "french" && to_learn_language == "farsi" {
    query = "SELECT woid, fran, pers, leid, imag FROM Words WHERE leid = ?"
  } else if user_language == "french" && to_learn_language == "german" {
    query = "SELECT woid, fran, germ, leid, imag FROM Words WHERE leid = ?"
  } else if user_language == "german" && to_learn_language == "french" {
    query = "SELECT woid, germ, fran, leid, imag FROM Words WHERE leid = ?"
  } else if user_language == "german" && to_learn_language == "farsi" {
    query = "SELECT woid, germ, pers, leid, imag FROM Words WHERE leid = ?"
  } else if user_language == "farsi" && to_learn_language == "french" {
    query = "SELECT woid, pers, fran, leid, imag FROM Words WHERE leid = ?"
  } else if user_language == "farsi" && to_learn_language == "german" {
    query = "SELECT woid, pers, germ, leid, imag FROM Words WHERE leid = ?"
  }
  return query
}


func numberLectures() ([]Lecture, error){
  // Find the number of lectures in the database

  var lectures []Lecture
  // Query to the database
  rows, err := db.Query("SELECT DISTINCT leid FROM Lecture")
  if err != nil {
    return nil, fmt.Errorf("numberLectures: %v", err)
  }

  defer rows.Close()
  // Store the results in the lecture object
  for rows.Next() {
    var lect Lecture
    if err := rows.Scan(&lect.ID); err != nil {
      return nil, fmt.Errorf("numberLectures  %v", err)
    }
    lectures = append(lectures, lect)
  }

  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("numberLectures %v", err)
  }

  return lectures, err
}



func wordsByWoid(Woid int, user_language string, to_learn_language string) (correctionWords) {
    // Get the words for
    var word Word
    // Get the words for this Woid
    var corrWor correctionWords
    // String giving the Query
    var query string

    // Query the database according to the scenario
    query = buildQueryWordByWoid(user_language, to_learn_language)
    rows, err := db.Query(query, Woid)
    if err != nil {
        return corrWor
    }

    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        if err := rows.Scan(&word.Woid, &word.Fran, &word.Pers, &word.Leid, &word.Imag); err != nil {
            return corrWor
        }
    }

    corrWor.QueryWord = word.Fran
    corrWor.CorrectWord = word.Pers

    return corrWor
}


func buildQueryWordByWoid(user_language string, to_learn_language string) (query string) {
  if user_language == "french" && to_learn_language == "farsi" {
    query = "SELECT woid, fran, pers, leid, imag FROM Words WHERE woid = ?"
  } else if user_language == "french" && to_learn_language == "german" {
    query = "SELECT woid, fran, germ, leid, imag FROM Words WHERE woid = ?"
  } else if user_language == "german" && to_learn_language == "french" {
    query = "SELECT woid, germ, fran, leid, imag FROM Words WHERE woid = ?"
  } else if user_language == "german" && to_learn_language == "farsi" {
    query = "SELECT woid, germ, pers, leid, imag FROM Words WHERE woid = ?"
  } else if user_language == "farsi" && to_learn_language == "french" {
    query = "SELECT woid, pers, fran, leid, imag FROM Words WHERE woid = ?"
  } else if user_language == "farsi" && to_learn_language == "german" {
    query = "SELECT woid, pers, germ, leid, imag FROM Words WHERE woid = ?"
  }
  return query
}
