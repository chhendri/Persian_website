package main

func defineTabWording (userLang string) (voc_title string, diag_title string, ex_image_title string, ex_word_title string) {
  // Define the title according to the language of the user
  if userLang == "french" {
    voc_title = "Vocabulaire"
    diag_title = "Dialogue"
    ex_image_title = "Exercice images"
    ex_word_title = "Exercice mots"

  } else if userLang == "english" {
    voc_title = "Vocabulary"
    diag_title = "Dialogue"
    ex_image_title = "Exercice images"
    ex_word_title = "Exercice words"

  } else if userLang == "dutch" {
    voc_title = "Woordenschat"
    diag_title = "Dialoog"
    ex_image_title = "Oefening beelden"
    ex_word_title = "Oefening woorden"

  } else if userLang == "german" {
    voc_title = "Wortschatz"
    diag_title = "Dialog"
    ex_image_title = "Übung Bilder"
    ex_word_title = "Übung Wörter"

  } else if userLang == "persian" {
    voc_title = "Vocabulary"
    diag_title = "Dialogue"
    ex_image_title = "Exercice images"
    ex_word_title = "Exercice words"
  }

  return voc_title, diag_title, ex_image_title, ex_word_title
}

func defineExerciceWording (userLang string) (corr_title string, ex1_ex2_title string, ex3_title string) {
  // Define the title according to the language of the user
  if userLang == "french" {
    corr_title = "Corriger"
    ex1_ex2_title = "Traduisez ces mots"
    ex3_title = "Combinez les traductions"

  } else if userLang == "english" {
    corr_title = "Correct"
    ex1_ex2_title = "Translate these words"
    ex3_title = "Combine the translations"

  } else if userLang == "dutch" {
    corr_title = "Verbeteren"
    ex1_ex2_title = "Vertaal deze woorden"
    ex3_title = "Combineer de vertalingen"

  } else if userLang == "german" {
    corr_title = "Korrigieren"
    ex1_ex2_title = "Übersetze diese Wörter"
    ex3_title = "Kombinier die Übersetzungen"

  } else if userLang == "persian" {
    corr_title = "Correct"
    ex1_ex2_title = "Translate these words"
    ex3_title = "Combine the translations"
  }

  return corr_title, ex1_ex2_title, ex3_title
}
