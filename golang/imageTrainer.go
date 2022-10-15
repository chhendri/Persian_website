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

func imageTrainer(leid string, Language_user string, Language_to_learn string) http.HandlerFunc {
  // Function to display the vocabulary of a lecture
    return func(w http.ResponseWriter, r *http.Request) {

      var Titles inputWords

      // Set the title of the page
      if Language_user == "farsi" {
        Titles.Title = "Learn vocabulary with images"
        Titles.SubTitle = "Hover over the image to get the translation"
      } else if Language_user == "french" {
        Titles.Title = "Apprenez le vocabulaire avec des images"
        Titles.SubTitle = "Passez votre souris sur une image pour avoir la traduction"
      } else if Language_user == "german" {
        Titles.Title = "Lerne Vokabular mit Bilder"
        Titles.SubTitle = "Geh mit ihrem Mause über dem Bild um die Überzetsung zu sehen"
      }

      tmpl_lec := template.Must(template.ParseFiles("html_files/" + Language_user + "/imageTrainerLecture_" + leid + "_" + Language_user + "_" + Language_to_learn + ".html"))
      tmpl_lec.Execute(w, Titles)
    }
}

func imageTrain_chooseLectureGet(user_language string) http.HandlerFunc {
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

    // Main function to display the imageTrain main page
    tmpl := template.Must(template.ParseFiles("html_files/imageTrainer.html"))
    tmpl.Execute(w, Titles)
  }
}

func imageTrain_chooseLecturePost(w http.ResponseWriter, r *http.Request){
  var buttonVal buttonValue

  r.ParseForm()
  for key := range r.PostForm {
    buttonVal.Val = key
  }

  http.Redirect(w, r, r.URL.Path + "/" + buttonVal.Val, http.StatusFound)
}


func getWordsImages(leid string, Language_user string, Language_to_learn string)(wordsImgs []wordsImagesHtml, err error){
  // Get all the words for a lecture

  // Query the database
  rows, err := db.Query(buildQueryImages(Language_user, Language_to_learn), leid)
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


func buildQueryImages(user_language string, to_learn_language string) (query string) {
  if user_language == "french" && to_learn_language == "farsi" {
    query = "SELECT b.NAME, a.fran, a.pers FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  } else if user_language == "french" && to_learn_language == "german" {
    query = "SELECT b.NAME, a.fran, a.germ FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  } else if user_language == "german" && to_learn_language == "french" {
    query = "SELECT b.NAME, a.germ, a.fran FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  } else if user_language == "german" && to_learn_language == "farsi" {
    query = "SELECT b.NAME, a.germ, a.pers FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  } else if user_language == "farsi" && to_learn_language == "french" {
    query = "SELECT b.NAME, a.pers, a.fran FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  } else if user_language == "farsi" && to_learn_language == "german" {
    query = "SELECT b.NAME, a.pers, a.germ FROM Words a RIGHT JOIN Images b ON a.Woid = b.WOID WHERE leid = ? AND imag = 'Y'"
  }
  return query
}
