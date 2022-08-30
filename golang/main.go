package main


import (
  "net/http"
  "github.com/gin-gonic/gin"
//  "log"
  "github.com/gorilla/mux"
  "html/template"
)

var templates *template.Template

func main()  {
  // Put the startPage as root path
  handleRequests()

  // Start server on localhost port 8000
  server := http.Server{
    Addr : ":8000",
    Handler : nil,
    ReadTimeout : 3000,
    WriteTimeout : 3000,
  }
  server.ListenAndServe()

}


func createHtml(w http.ResponseWriter, r *http.Request) {
  connectDB(w, r)
  // Construct the page with the list of lectures
  constr_htmlWordlist("listWords")
  constr_htmlWordlist("imageTrainer")
  constr_htmlWordlist("wordTrainer")
  createHtmlImages_allLectures()
  router := gin.Default()
  // Display the words per lecture
  router.GET("/listWords/:id", getwordsByLecture)
  router.Run("localhost:8080")
}


func handleRequests() {
    // Handler of paths and associated html pages

    myRouter := mux.NewRouter()

    // Main page
    templates = template.Must(template.ParseGlob("html_files/*.html"))

    myRouter.HandleFunc("/", startPageGet).Methods("GET")
    myRouter.HandleFunc("/", startPagePost).Methods("POST")
    http.Handle("/", myRouter)

    // Training with images
    myRouter.HandleFunc("/imageTrain", imageTrain_chooseLectureGet).Methods("GET")
    myRouter.HandleFunc("/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/imageTrain", myRouter)
    myRouter.HandleFunc("/imageTrain/lecture1", imageTrainer("1"))
    myRouter.HandleFunc("/imageTrain/lecture2", imageTrainer("2"))
    myRouter.HandleFunc("/imageTrain/lecture3", imageTrainer("3"))
    myRouter.HandleFunc("/imageTrain/lecture4", imageTrainer("4"))
    myRouter.HandleFunc("/imageTrain/lecture5", imageTrainer("5"))
    myRouter.HandleFunc("/imageTrain/lecture6", imageTrainer("6"))
    myRouter.HandleFunc("/imageTrain/lecture7", imageTrainer("7"))
    myRouter.HandleFunc("/imageTrain/lecture8", imageTrainer("8"))
    myRouter.HandleFunc("/imageTrain/lecture9", imageTrainer("9"))
    myRouter.HandleFunc("/imageTrain/lecture10", imageTrainer("10"))
    myRouter.HandleFunc("/imageTrain/lecture11", imageTrainer("11"))

    // Training with words
    myRouter.HandleFunc("/wordTrain", wordTrainerGet).Methods("GET")
    myRouter.HandleFunc("/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/wordTrain", myRouter)
    myRouter.HandleFunc("/wordTrain/lecture1", askWordVisual)
    //myRouter.HandleFunc("/wordTrain/trainLecture1", lectureTrainer)

    // List of the words per lecture
    myRouter.HandleFunc("/listWords", listWordsGet).Methods("GET")
    myRouter.HandleFunc("/listWords", listWordsPost).Methods("POST")
    http.Handle("/listWords", myRouter)
    myRouter.HandleFunc("/listWords/lecture1", lectureHandler("1"))
    myRouter.HandleFunc("/listWords/lecture2", lectureHandler("2"))
    myRouter.HandleFunc("/listWords/lecture3", lectureHandler("3"))
    myRouter.HandleFunc("/listWords/lecture4", lectureHandler("4"))
    myRouter.HandleFunc("/listWords/lecture5", lectureHandler("5"))
    myRouter.HandleFunc("/listWords/lecture6", lectureHandler("6"))
    myRouter.HandleFunc("/listWords/lecture7", lectureHandler("7"))
    myRouter.HandleFunc("/listWords/lecture8", lectureHandler("8"))
    myRouter.HandleFunc("/listWords/lecture9", lectureHandler("9"))
    myRouter.HandleFunc("/listWords/lecture10", lectureHandler("10"))
    myRouter.HandleFunc("/listWords/lecture11", lectureHandler("11"))

    // Creation of the HTML files
    myRouter.HandleFunc("/createHtml", createHtml)

    http.ListenAndServe(":8000", myRouter)
}
