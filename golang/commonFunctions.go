package main


import (
  "log"
  "io/ioutil"
  "net/http"
  "strconv"
  "fmt"
)


func goHome(w http.ResponseWriter, r *http.Request){
  http.Redirect(w, r, "/", http.StatusFound)
}



func constr_htmlLectures (headerFileName string) (){
  // Construct the html file for listWords

  var htmlString string
  // Find the number of lectures in the 
  base
  n_lectures, err := numberLectures()
  // Read the header file
  htmlHeaderFile := headerFileName + "_header.html"
  header, err := ioutil.ReadFile("html_files/" + htmlHeaderFile)
  if err != nil {
      log.Fatal(err)
  }
  htmlString += string(header)

  htmlString += "<h1>{{ .Title }}</h1>\n"

  htmlString += "<p>\n<div align='center'>\n<form method='POST'>\n"
  // Add a button for each lecture
  for _ , i := range n_lectures {
    lecture_id := "lecture" + strconv.Itoa(i.ID)
    lecture_text := "{{ .SubTitle }} " + strconv.Itoa(i.ID)
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


func constr_htmlWordList () {
  // Find the number of lectures in the database
  n_lectures, err := numberLectures()
  fmt.Println(err)
  for _ , i := range n_lectures {
      id := strconv.Itoa(i.ID)
      tableVocabulary(wordsByLecture(id, "french", "farsi"), id, "french", "farsi")
      tableVocabulary(wordsByLecture(id, "german", "farsi"), id, "german", "farsi")
      tableVocabulary(wordsByLecture(id, "french", "german"), id, "french", "german")
      tableVocabulary(wordsByLecture(id, "german", "french"), id, "german", "french")
      tableVocabulary(wordsByLecture(id, "farsi", "french"), id, "farsi", "french")
      tableVocabulary(wordsByLecture(id, "farsi", "german"), id, "farsi", "german")
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
    createHtmlImages("imageTrainerLecture", leid, "french", "farsi")
    createHtmlImages("imageTrainerLecture", leid, "french", "german")
    createHtmlImages("imageTrainerLecture", leid, "farsi", "french")
    createHtmlImages("imageTrainerLecture", leid, "farsi", "german")
    createHtmlImages("imageTrainerLecture", leid, "german", "farsi")
    createHtmlImages("imageTrainerLecture", leid, "german", "french")
  }
}


func createHtmlWordTrain_allLectures() {
  var leid string
  n_lectures, err := numberLectures()
  if err != nil {
    log.Fatal(err)
  }
  for _ , i := range n_lectures {
    leid = strconv.Itoa(i.ID)
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "french", "farsi")
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "french", "german")
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "farsi", "french")
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "farsi", "german")
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "german", "farsi")
    constrHtmlWordTrainLecture("wordTrainerLecture", leid, "german", "french")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "french", "farsi")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "french", "german")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "farsi", "french")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "farsi", "german")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "german", "farsi")
    constrHtmlWordTrainLectureCorrection("wordTrainerCorrLecture", leid, "german", "french")
  }
}


func createHtmlImages (headerFileName string, leid string, Language_user string, Language_to_learn string) {
  // Find the words, translations and image paths
  wordsImgs, err := getWordsImages(leid, Language_user, Language_to_learn)
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
    htmlString += "<img src='../../Data/images/" + i.Img_path + ".jpeg' style='width:300px;height:200px'>\n"
    htmlString += "</div>\n<div class='flip-box-back'>\n"
    htmlString += "<h2>" + i.Pers + "</h2>\n"
    htmlString += "<h2>" + i.Trans + "</h2>\n"
    htmlString += "</div>\n</div>\n</div>\n"
  }
  // Add the footer
  htmlString += "</body>\n</html>"

  // Save to file
  htmlFile := headerFileName + "_" + leid + "_" + Language_user + "_" + Language_to_learn + ".html"
  if err := ioutil.WriteFile("html_files/" + Language_user + "/" + htmlFile, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }
}


func constrHtmlWordTrainLecture (headerFileName string, leid string, user_language string, to_learn_language string) {
  // Find the words and translations
  wordsLect := wordsByLecture(leid, user_language, to_learn_language)

  // Construct the html file for listWords
  var htmlString string
  // Read the header file
  htmlHeaderFile := headerFileName + "_header.html"
  header, err := ioutil.ReadFile("html_files/" + htmlHeaderFile)
  if err != nil {
      log.Fatal(err)
  }
  htmlString += string(header)

  htmlString += "</div>\n</header>\n<body><div class='container'>\n"
  htmlString += "<form id='survey-form' action='/" + user_language + "To" + to_learn_language + "/wordTrain/lecture" + leid + "/corr' method='POST'>"

  for _ , i := range wordsLect {
    htmlString += "<div class='labels'>\n"
    htmlString += "<label id='name-label' for='word_to_translate'>" + i.Fran + "</label>\n"
    htmlString += "</div>\n"
    htmlString += "<div class='input-tab'>\n"
    htmlString += "<input class='input-field' type='text' id='response_" + strconv.Itoa(i.Woid) + "' name='response_" + strconv.Itoa(i.Woid) + "' placeholder='Enter your response' autofocus>\n"
    htmlString += "</div>\n"
  }

  // Add the footer
  htmlString += "<div class='btn'>\n<button id='submit' type='submit' name='submit' value='submit'>Submit</button>\n</div>\n"
  htmlString += "</form>\n</div>\n</body>\n</html>\n"

  // Save to file
  htmlFile := headerFileName + "_" + leid + "_" + user_language + "_" + to_learn_language + ".html"
  if err := ioutil.WriteFile("html_files/" + user_language + "/" + htmlFile, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }
}



func constrHtmlWordTrainLectureCorrection (headerFileName string, leid string, user_language string, to_learn_language string) (){

  // Construct the html file for listWords
  var htmlString string

  // Read the header file
  htmlHeaderFile := headerFileName + "_header.html"
  header, err := ioutil.ReadFile("html_files/" + htmlHeaderFile)
  if err != nil {
      log.Fatal(err)
  }
  htmlString += string(header)
  htmlString += "</div>\n</header>\n<body><div class='container'>\n"
  htmlString += "{{ range .Words }}\n"
  htmlString += "<form id='survey-form' action='/" + user_language + "To" + to_learn_language + "/wordTrain/lecture" + leid + "/corr' method='POST'>\n"
  htmlString += "<div class='input-tab'>\n"
  htmlString += "<label class='column-3 left' id='name-label' for='word_to_translate'>{{ .QueryWord }}</label>\n"
  htmlString += "<font color='{{ .Color }}'><label class='column-3 center' id='response' name='response'>{{ .UserWord }}</label></font>\n"
  htmlString += "<label class='column-3 right' id='correct' name='correct'>{{ .CorrectWord }}</label>\n"
  htmlString += "</div>\n"
  htmlString += "{{ end }}\n"
  // Add the footer
  htmlString += "</form>\n</div>\n</body>\n</html>\n"

  // Save to file
  htmlFile := headerFileName + "_" + leid + "_" + user_language + "_" + to_learn_language + ".html"
  if err := ioutil.WriteFile("html_files/" + user_language + "/" + htmlFile, []byte(htmlString), 0666); err != nil {
    log.Fatal(err)
  }

}
