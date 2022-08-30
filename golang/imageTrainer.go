package main


import (
  "fmt"
  "net/http"
  "html/template"
)

type wordsImagesHtml struct {
  Pers      string   `json:"pers"`
  Trans     string   `json:"trans"`
  Img_path  string   `json:"img_path"`
}

func imageTrainer(leid string) http.HandlerFunc {
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {
      tmpl_lec := template.Must(template.ParseFiles("html_files/imageTrainerLecture_" + leid + ".html"))
      tmpl_lec.Execute(w, nil)
    }
}


func imageTrain_chooseLectureGet (w http.ResponseWriter, r *http.Request) {
    // Main function to display the imageTrain main page
    tmpl := template.Must(template.ParseFiles("html_files/imageTrainer.html"))
    tmpl.Execute(w, nil)
}

func imageTrain_chooseLecturePost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  http.Redirect(w, r, "/imageTrain/" + buttonVal.Val, http.StatusFound)
}


func getWordsImages(leid string)(wordsImgs []wordsImagesHtml, err error){
  // Get all the words for a lecture

  // Query the database
  rows, err := db.Query("SELECT b.NAME, a.fran, a.pers FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'", leid)
  if err != nil {
      return nil, fmt.Errorf("getWordsImages %q: %v", leid, err)
  }

  defer rows.Close()
  // Loop through rows, using Scan to assign column data to struct fields.
  for rows.Next() {
      var wor wordsImagesHtml
      if err := rows.Scan(&wor.Img_path, &wor.Trans, &wor.Pers); err != nil {
          return nil, fmt.Errorf("getWordsImages %q: %v", leid, err)
      }
      wordsImgs = append(wordsImgs, wor)
  }
  if err := rows.Err(); err != nil {
      return nil, fmt.Errorf("getWordsImages %q: %v", leid, err)
  }

  return wordsImgs, nil
}
