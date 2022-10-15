package main


import (
  "net/http"
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
  constr_htmlLectures("listWords")
  constr_htmlLectures("imageTrainer")
  constr_htmlLectures("wordTrainer")
  createHtmlImages_allLectures()
  createHtmlWordTrain_allLectures()
  constr_htmlWordList()
}


func handleRequests() {
    // Handler of paths and associated html pages

    myRouter := mux.NewRouter()

    // Main page
    templates = template.Must(template.ParseGlob("html_files/*/*.html"))


    // Handler for the start page
    myRouter.HandleFunc("/", startPageGet).Methods("GET")
    myRouter.HandleFunc("/", startPagePost).Methods("POST")
    http.Handle("/", myRouter)
    // Handler for the language of the user
    myRouter.HandleFunc("/chooseLanguageUser", startPageLangPost).Methods("POST")
    http.Handle("/chooseLanguageUser", myRouter)
    // Handler for the language to learn
    myRouter.HandleFunc("/chooseLanguageToLearn", startPageLangToLearnPost).Methods("POST")
    http.Handle("/chooseLanguageToLearn", myRouter)
    // Handler for the action
    myRouter.HandleFunc("/chooseAction", startPageActionPost).Methods("POST")
    http.Handle("/chooseAction", myRouter)

    // Get back to home page
    myRouter.HandleFunc("/goHome", goHome)
    http.Handle("/goHome", myRouter)

    // Training with images

    // Farsi to French
    myRouter.HandleFunc("/farsiTofrench/imageTrain", imageTrain_chooseLectureGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/farsiTofrench/imageTrain", myRouter)
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture1", imageTrainer("1", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture2", imageTrainer("2", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture3", imageTrainer("3", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture4", imageTrainer("4", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture5", imageTrainer("5", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture6", imageTrainer("6", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture7", imageTrainer("7", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture8", imageTrainer("8", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture9", imageTrainer("9", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture10", imageTrainer("10", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/imageTrain/lecture11", imageTrainer("11", "farsi", "french"))

    // Farsi to german
    myRouter.HandleFunc("/farsiTogerman/imageTrain", imageTrain_chooseLectureGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/farsiTogerman/imageTrain", myRouter)
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture1", imageTrainer("1", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture2", imageTrainer("2", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture3", imageTrainer("3", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture4", imageTrainer("4", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture5", imageTrainer("5", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture6", imageTrainer("6", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture7", imageTrainer("7", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture8", imageTrainer("8", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture9", imageTrainer("9", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture10", imageTrainer("10", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/imageTrain/lecture11", imageTrainer("11", "farsi", "german"))

    // french to german
    myRouter.HandleFunc("/frenchTogerman/imageTrain", imageTrain_chooseLectureGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/frenchTogerman/imageTrain", myRouter)
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture1", imageTrainer("1", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture2", imageTrainer("2", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture3", imageTrainer("3", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture4", imageTrainer("4", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture5", imageTrainer("5", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture6", imageTrainer("6", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture7", imageTrainer("7", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture8", imageTrainer("8", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture9", imageTrainer("9", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture10", imageTrainer("10", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/imageTrain/lecture11", imageTrainer("11", "french", "german"))

    // french to farsi
    myRouter.HandleFunc("/frenchTofarsi/imageTrain", imageTrain_chooseLectureGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/frenchTofarsi/imageTrain", myRouter)
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture1", imageTrainer("1", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture2", imageTrainer("2", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture3", imageTrainer("3", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture4", imageTrainer("4", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture5", imageTrainer("5", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture6", imageTrainer("6", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture7", imageTrainer("7", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture8", imageTrainer("8", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture9", imageTrainer("9", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture10", imageTrainer("10", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/imageTrain/lecture11", imageTrainer("11", "french", "farsi"))

    // german to farsi
    myRouter.HandleFunc("/germanTofarsi/imageTrain", imageTrain_chooseLectureGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/germanTofarsi/imageTrain", myRouter)
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture1", imageTrainer("1", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture2", imageTrainer("2", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture3", imageTrainer("3", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture4", imageTrainer("4", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture5", imageTrainer("5", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture6", imageTrainer("6", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture7", imageTrainer("7", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture8", imageTrainer("8", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture9", imageTrainer("9", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture10", imageTrainer("10", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/imageTrain/lecture11", imageTrainer("11", "german", "farsi"))

    // german to french
    myRouter.HandleFunc("/germanTofrench/imageTrain", imageTrain_chooseLectureGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/imageTrain", imageTrain_chooseLecturePost).Methods("POST")
    http.Handle("/germanTofrench/imageTrain", myRouter)
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture1", imageTrainer("1", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture2", imageTrainer("2", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture3", imageTrainer("3", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture4", imageTrainer("4", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture5", imageTrainer("5", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture6", imageTrainer("6", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture7", imageTrainer("7", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture8", imageTrainer("8", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture9", imageTrainer("9", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture10", imageTrainer("10", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/imageTrain/lecture11", imageTrainer("11", "german", "french"))




    // Training with words

    // french to farsi
    myRouter.HandleFunc("/frenchTofarsi/wordTrain", wordTrainerGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/frenchTofarsi/wordTrain", myRouter)
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture1", wordTrainLectureGet("1", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture2", wordTrainLectureGet("2", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture3", wordTrainLectureGet("3", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture4", wordTrainLectureGet("4", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture5", wordTrainLectureGet("5", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture6", wordTrainLectureGet("6", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture7", wordTrainLectureGet("7", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture8", wordTrainLectureGet("8", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture9", wordTrainLectureGet("9", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture10", wordTrainLectureGet("10", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture11", wordTrainLectureGet("11", "french", "farsi")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture1/corr", wordTrainLecturePost("1", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture2/corr", wordTrainLecturePost("2", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture3/corr", wordTrainLecturePost("3", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture4/corr", wordTrainLecturePost("4", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture5/corr", wordTrainLecturePost("5", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture6/corr", wordTrainLecturePost("6", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture7/corr", wordTrainLecturePost("7", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture8/corr", wordTrainLecturePost("8", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture9/corr", wordTrainLecturePost("9", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture10/corr", wordTrainLecturePost("10", "french", "farsi")).Methods("POST")
    myRouter.HandleFunc("/frenchTofarsi/wordTrain/lecture11/corr", wordTrainLecturePost("11", "french", "farsi")).Methods("POST")


    // french to german
    myRouter.HandleFunc("/frenchTogerman/wordTrain", wordTrainerGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/frenchTogerman/wordTrain", myRouter)
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture1", wordTrainLectureGet("1", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture2", wordTrainLectureGet("2", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture3", wordTrainLectureGet("3", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture4", wordTrainLectureGet("4", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture5", wordTrainLectureGet("5", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture6", wordTrainLectureGet("6", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture7", wordTrainLectureGet("7", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture8", wordTrainLectureGet("8", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture9", wordTrainLectureGet("9", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture10", wordTrainLectureGet("10", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture11", wordTrainLectureGet("11", "french", "german")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture1/corr", wordTrainLecturePost("1", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture2/corr", wordTrainLecturePost("2", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture3/corr", wordTrainLecturePost("3", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture4/corr", wordTrainLecturePost("4", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture5/corr", wordTrainLecturePost("5", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture6/corr", wordTrainLecturePost("6", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture7/corr", wordTrainLecturePost("7", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture8/corr", wordTrainLecturePost("8", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture9/corr", wordTrainLecturePost("9", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture10/corr", wordTrainLecturePost("10", "french", "german")).Methods("POST")
    myRouter.HandleFunc("/frenchTogerman/wordTrain/lecture11/corr", wordTrainLecturePost("11", "french", "german")).Methods("POST")


    // farsi to german
    myRouter.HandleFunc("/farsiTogerman/wordTrain", wordTrainerGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/farsiTogerman/wordTrain", myRouter)
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture1", wordTrainLectureGet("1", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture2", wordTrainLectureGet("2", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture3", wordTrainLectureGet("3", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture4", wordTrainLectureGet("4", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture5", wordTrainLectureGet("5", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture6", wordTrainLectureGet("6", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture7", wordTrainLectureGet("7", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture8", wordTrainLectureGet("8", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture9", wordTrainLectureGet("9", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture10", wordTrainLectureGet("10", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture11", wordTrainLectureGet("11", "farsi", "german")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture1/corr", wordTrainLecturePost("1", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture2/corr", wordTrainLecturePost("2", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture3/corr", wordTrainLecturePost("3", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture4/corr", wordTrainLecturePost("4", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture5/corr", wordTrainLecturePost("5", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture6/corr", wordTrainLecturePost("6", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture7/corr", wordTrainLecturePost("7", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture8/corr", wordTrainLecturePost("8", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture9/corr", wordTrainLecturePost("9", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture10/corr", wordTrainLecturePost("10", "farsi", "german")).Methods("POST")
    myRouter.HandleFunc("/farsiTogerman/wordTrain/lecture11/corr", wordTrainLecturePost("11", "farsi", "german")).Methods("POST")


    // farsi to french
    myRouter.HandleFunc("/farsiTofrench/wordTrain", wordTrainerGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/farsiTofrench/wordTrain", myRouter)
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture1", wordTrainLectureGet("1", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture2", wordTrainLectureGet("2", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture3", wordTrainLectureGet("3", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture4", wordTrainLectureGet("4", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture5", wordTrainLectureGet("5", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture6", wordTrainLectureGet("6", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture7", wordTrainLectureGet("7", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture8", wordTrainLectureGet("8", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture9", wordTrainLectureGet("9", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture10", wordTrainLectureGet("10", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture11", wordTrainLectureGet("11", "farsi", "french")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture1/corr", wordTrainLecturePost("1", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture2/corr", wordTrainLecturePost("2", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture3/corr", wordTrainLecturePost("3", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture4/corr", wordTrainLecturePost("4", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture5/corr", wordTrainLecturePost("5", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture6/corr", wordTrainLecturePost("6", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture7/corr", wordTrainLecturePost("7", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture8/corr", wordTrainLecturePost("8", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture9/corr", wordTrainLecturePost("9", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture10/corr", wordTrainLecturePost("10", "farsi", "french")).Methods("POST")
    myRouter.HandleFunc("/farsiTofrench/wordTrain/lecture11/corr", wordTrainLecturePost("11", "farsi", "french")).Methods("POST")


    // german to french
    myRouter.HandleFunc("/germanTofrench/wordTrain", wordTrainerGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/germanTofrench/wordTrain", myRouter)
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture1", wordTrainLectureGet("1", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture2", wordTrainLectureGet("2", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture3", wordTrainLectureGet("3", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture4", wordTrainLectureGet("4", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture6", wordTrainLectureGet("6", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture5", wordTrainLectureGet("5", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture7", wordTrainLectureGet("7", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture8", wordTrainLectureGet("8", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture9", wordTrainLectureGet("9", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture10", wordTrainLectureGet("10", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture11", wordTrainLectureGet("11", "german", "french")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture1/corr", wordTrainLecturePost("1", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture2/corr", wordTrainLecturePost("2", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture3/corr", wordTrainLecturePost("3", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture4/corr", wordTrainLecturePost("4", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture5/corr", wordTrainLecturePost("5", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture6/corr", wordTrainLecturePost("6", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture7/corr", wordTrainLecturePost("7", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture8/corr", wordTrainLecturePost("8", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture9/corr", wordTrainLecturePost("9", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture10/corr", wordTrainLecturePost("10", "german", "french")).Methods("POST")
    myRouter.HandleFunc("/germanTofrench/wordTrain/lecture11/corr", wordTrainLecturePost("11", "german", "french")).Methods("POST")


    // german to farsi
    myRouter.HandleFunc("/germanTofarsi/wordTrain", wordTrainerGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain", wordTrainerPost).Methods("POST")
    http.Handle("/germanTofarsi/wordTrain", myRouter)
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture1", wordTrainLectureGet("1", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture2", wordTrainLectureGet("2", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture3", wordTrainLectureGet("3", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture4", wordTrainLectureGet("4", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture5", wordTrainLectureGet("5", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture6", wordTrainLectureGet("6", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture8", wordTrainLectureGet("8", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture7", wordTrainLectureGet("7", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture9", wordTrainLectureGet("9", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture10", wordTrainLectureGet("10", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture11", wordTrainLectureGet("11", "german", "farsi")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture1/corr", wordTrainLecturePost("1", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture2/corr", wordTrainLecturePost("2", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture3/corr", wordTrainLecturePost("3", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture4/corr", wordTrainLecturePost("4", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture5/corr", wordTrainLecturePost("5", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture6/corr", wordTrainLecturePost("6", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture7/corr", wordTrainLecturePost("7", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture8/corr", wordTrainLecturePost("8", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture9/corr", wordTrainLecturePost("9", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture10/corr", wordTrainLecturePost("10", "german", "farsi")).Methods("POST")
    myRouter.HandleFunc("/germanTofarsi/wordTrain/lecture11/corr", wordTrainLecturePost("11", "german", "farsi")).Methods("POST")




    // List of the words per lecture

    // french to farsi
    myRouter.HandleFunc("/frenchTofarsi/listWords", listWordsGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTofarsi/listWords", listWordsPost).Methods("POST")
    http.Handle("/frenchTofarsi/listWords", myRouter)
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture1", lectureHandler("1", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture2", lectureHandler("2", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture3", lectureHandler("3", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture4", lectureHandler("4", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture5", lectureHandler("5", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture6", lectureHandler("6", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture7", lectureHandler("7", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture8", lectureHandler("8", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture9", lectureHandler("9", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture10", lectureHandler("10", "french", "farsi"))
    myRouter.HandleFunc("/frenchTofarsi/listWords/lecture11", lectureHandler("11", "french", "farsi"))

    // french to german
    myRouter.HandleFunc("/frenchTogerman/listWords", listWordsGet("french")).Methods("GET")
    myRouter.HandleFunc("/frenchTogerman/listWords", listWordsPost).Methods("POST")
    http.Handle("/frenchTogerman/listWords", myRouter)
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture1", lectureHandler("1", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture2", lectureHandler("2", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture3", lectureHandler("3", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture4", lectureHandler("4", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture5", lectureHandler("5", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture6", lectureHandler("6", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture7", lectureHandler("7", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture8", lectureHandler("8", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture9", lectureHandler("9", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture10", lectureHandler("10", "french", "german"))
    myRouter.HandleFunc("/frenchTogerman/listWords/lecture11", lectureHandler("11", "french", "german"))

    // farsi to german
    myRouter.HandleFunc("/farsiTogerman/listWords", listWordsGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTogerman/listWords", listWordsPost).Methods("POST")
    http.Handle("/farsiTogerman/listWords", myRouter)
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture1", lectureHandler("1", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture2", lectureHandler("2", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture3", lectureHandler("3", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture4", lectureHandler("4", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture5", lectureHandler("5", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture6", lectureHandler("6", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture7", lectureHandler("7", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture8", lectureHandler("8", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture9", lectureHandler("9", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture10", lectureHandler("10", "farsi", "german"))
    myRouter.HandleFunc("/farsiTogerman/listWords/lecture11", lectureHandler("11", "farsi", "german"))

    // farsi to french
    myRouter.HandleFunc("/farsiTofrench/listWords", listWordsGet("farsi")).Methods("GET")
    myRouter.HandleFunc("/farsiTofrench/listWords", listWordsPost).Methods("POST")
    http.Handle("/farsiTofrench/listWords", myRouter)
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture1", lectureHandler("1", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture2", lectureHandler("2", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture3", lectureHandler("3", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture4", lectureHandler("4", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture5", lectureHandler("5", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture6", lectureHandler("6", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture7", lectureHandler("7", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture8", lectureHandler("8", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture9", lectureHandler("9", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture10", lectureHandler("10", "farsi", "french"))
    myRouter.HandleFunc("/farsiTofrench/listWords/lecture11", lectureHandler("11", "farsi", "french"))

    // german to french
    myRouter.HandleFunc("/germanTofrench/listWords", listWordsGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofrench/listWords", listWordsPost).Methods("POST")
    http.Handle("/germanTofrench/listWords", myRouter)
    myRouter.HandleFunc("/germanTofrench/listWords/lecture1", lectureHandler("1", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture2", lectureHandler("2", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture3", lectureHandler("3", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture4", lectureHandler("4", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture5", lectureHandler("5", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture6", lectureHandler("6", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture7", lectureHandler("7", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture8", lectureHandler("8", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture9", lectureHandler("9", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture10", lectureHandler("10", "german", "french"))
    myRouter.HandleFunc("/germanTofrench/listWords/lecture11", lectureHandler("11", "german", "french"))

    // german to farsi
    myRouter.HandleFunc("/germanTofarsi/listWords", listWordsGet("german")).Methods("GET")
    myRouter.HandleFunc("/germanTofarsi/listWords", listWordsPost).Methods("POST")
    http.Handle("/germanTofarsi/listWords", myRouter)
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture1", lectureHandler("1", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture2", lectureHandler("2", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture3", lectureHandler("3", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture4", lectureHandler("4", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture5", lectureHandler("5", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture6", lectureHandler("6", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture7", lectureHandler("7", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture8", lectureHandler("8", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture9", lectureHandler("9", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture10", lectureHandler("10", "german", "farsi"))
    myRouter.HandleFunc("/germanTofarsi/listWords/lecture11", lectureHandler("11", "german", "farsi"))

    // Creation of the HTML files
    myRouter.HandleFunc("/createHtml", createHtml)

    http.ListenAndServe(":8000", myRouter)
}
