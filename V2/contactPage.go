package main

import (
  "net/http"
  "strings"
)


func renderContactPage(w http.ResponseWriter, r *http.Request) {
  // Choose the user language
  userLang := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/lang/"), "/contact")
  // Set the header and the text of the start page
  header := headerWording(userLang)
  title, text_1, text_2, text_3, text_4 := wordingContactPage(userLang)
  data := struct {
    Header   Header
    Title    string
    Text_1   string
    Text_2   string
    Text_3   string
    Text_4   string
  }{
    header,
    title,
    text_1,
    text_2,
    text_3,
    text_4,
  }
  templates.ExecuteTemplate(w, "contact_page.html", data)
}

func wordingContactPage (userLang string) (title string, text_1 string, text_2 string, text_3 string, text_4 string) {
  if userLang == "french" {
    title = "À propos de nous"
    text_1 = "Ce site veut fournir un outil d'étude de vocabulaire. Son but principal est de fournir une ressource accessible partout, à chaque moment et gratuitement."
    text_2 = "Ce site a été créé par Charlotte Hendrickx. Charlotte parle couremment Français, Néerlandais et Anglais et voulait apprendre le Persan. Elle trouvait cela difficile de trouver des ressources pour apprendre le Persan. Ces ressources sont souvent disponibles pour d'autres langues, mais pas pour celle-ci."
    text_3 = "Une aide précieuse a été apportée par Nehmat Husseini, qui parle Persan et Allemand et a aidé avec les traductions vers ces deux langues. Nous remercions aussi Chris Adam, qui a aidé pour la logistique de cette application."
    text_4 = "Vu qu'aucun des deux créateurs de ce site sont des linguistes professionnels, nous nous excusons pour toute erreur dans le vocabulaire présenté ici. "
  } else if userLang == "english" {
    title = "About us"
    text_1 = "This site wants to provide a vocabulary playground. Its' main aim is to have a learning resource that is available anywhere, anytime and for free. "
    text_2 = "This site was created by Charlotte Hendrickx. Charlotte speaks French, Dutch and English fluently and wanted to learn Persian. She found it difficult to find free resources to learn Persian. These resources are often found for other languages, but not for Persian. "
    text_3 = "A precious help was provided by Nehmat Husseini, who speaks Persian and German and helped with the translation towards these two languages. Another big thank you goes to Chris Adam who helped with the logistics of this app."
    text_4 = "As neither of the creators of the site are professional languists, we apologize for any error in the vocabulary on this site."
  } else if userLang == "dutch" {
    title = "Over ons"
    text_1 = "Deze website werdt aangemaakt om woordenschat altijd, overal en gratis te kunnen leren. "
    text_2 = "Deze website werdt door Charlotte Hendrickx gecrëerd. Charlotte spreekt vloeiend Frans, Nederlands en Engels en wou Persisch leren. Ze vondt het moeilijk om gratis materiaal te vinden om Persisch te leren, ook al zijn deze voor andere talen beschikbaar."
    text_3 = "Een waardevolle hulp werd door Nehmat Husseini geboden. Deze spreekt Persisch en Duits en hielp met de vertaling naar deze beide talen. Ook Chris Adam hielp met de logistiek van deze Webapp en wordt bedankt."
    text_4 = "Aangezien geen van de makers van deze site professionele taalkundigen zijn, verontschuldigen wij ons voor eventuele fouten in de woordenschat op deze site."
  } else if userLang == "german" {
    title = "Über uns"
    text_1 = "Diese Website will einen Spielplatz für Vokabeln bieten. Ihr Hauptziel ist es, eine Lernressource bereitzustellen, die überall, jederzeit und kostenlos verfügbar ist. "
    text_2 = "Diese Website wurde von Charlotte Hendrickx erstellt. Charlotte spricht fließend Französisch, Niederländisch und Englisch und wollte Persisch lernen. Sie fand es schwierig, kostenlose Ressourcen zum Erlernen der persischen Sprache zu finden. Diese Ressourcen gibt es oft für andere Sprachen, aber nicht für Persisch. "
    text_3 = "Eine wertvolle Hilfe war Nehmat Husseini, die Persisch und Deutsch spricht und bei der Übersetzung in diese beiden Sprachen geholfen hat. Vielen Dank auch an Chris Adam, der mit dem Logistik von dieses App geholfen hat."
    text_4 = "Da keiner der Ersteller dieser Website ein professioneller Sprachwissenschaftler ist, entschuldigen wir uns für eventuelle Fehler im Vokabular auf dieser Website."
  } else if userLang == "persian" {
    title = "About us"
    text_1 = "This site wants to provide a vocabulary playground. Its' main aim is to have a learning resource that is available anywhere, anytime and for free. "
    text_2 = "This site was created by Charlotte Hendrickx. Charlotte speaks French, Dutch and English fluently and wanted to learn Persian. She found it difficult to find free resources to learn Persian. These resources are often found for other languages, but not for Persian. "
    text_3 = "A precious help was provided by Nehmat Husseini, who speaks Persian and German and helped with the translation towards these two languages."
    text_4 = "As neither of the creators of the site are professional languists, we apologize for any error in the vocabulary on this site."
  }
  return title, text_1, text_2, text_3, text_4
}
