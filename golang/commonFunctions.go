package main


import (
  "log"
  "io/ioutil"
  "strconv"
)


func constr_htmlWordlist (headerFileName string) (){
  // Construct the html file for listWords

  var htmlString string
  // Find the number of lectures in the database
  n_lectures, err := numberLectures()
  // Read the header file
  htmlHeaderFile := headerFileName + "_header.html"
  header, err := ioutil.ReadFile("html_files/" + htmlHeaderFile)
  if err != nil {
      log.Fatal(err)
  }
  htmlString += string(header)
  htmlString += "<p>\n<div align='center'>\n<form method='POST'>\n"
  // Add a button for each lecture
  for _ , i := range n_lectures {
    lecture_id := "lecture" + strconv.Itoa(i.ID)
    lecture_text := "Lecture " + strconv.Itoa(i.ID)
    htmlString += "<button type='submit' value='" + lecture_id + "' name='" + lecture_id + "'>" + lecture_text + "</button>\n"
  }
  // Add the footer
  htmlString += "</div>\n</form>\n</p>\n"
  htmlString += "</body>\n</html>"

  // Save to file
  htmlFile := headerFileName + ".html"
  if err := ioutil.WriteFile("html_files/" + htmlFile, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }
}



func createHtmlImages_allLectures() {
  var leid string
  n_lectures, err := numberLectures()
  if err != nil {
    log.Fatal(err)
  }
  for _ , i := range n_lectures {
    leid = strconv.Itoa(i.ID)
    createHtmlImages("imageTrainerLecture", leid)
  }
}


func createHtmlImages (headerFileName string, leid string) {
  // Find the words, translations and image paths
  wordsImgs, err := getWordsImages(leid)
  // Construct the html file for listWords
  var htmlString string
  // Read the header file
  htmlHeaderFile := headerFileName + "_header.html"
  header, err := ioutil.ReadFile("html_files/" + htmlHeaderFile)
  if err != nil {
      log.Fatal(err)
  }
  htmlString += string(header)
  // Add a button for each lecture
  for _ , i := range wordsImgs {
    htmlString += "<div class='flip-box'>\n<div class='flip-box-inner'>\n<div class='flip-box-front'>\n"
    htmlString += "<img src='../Data/images/" + i.Img_path + ".jpeg' style='width:300px;height:200px'>\n"
    htmlString += "</div>\n<div class='flip-box-back'>\n"
    htmlString += "<h2>" + i.Pers + "</h2>\n"
    htmlString += "<h2>" + i.Trans + "</h2>\n"
    htmlString += "</div>\n</div>\n</div>\n"
  }
  // Add the footer
  htmlString += "</body>\n</html>"

  // Save to file
  htmlFile := headerFileName + "_" + leid + ".html"
  if err := ioutil.WriteFile("html_files/" + htmlFile, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }
}
