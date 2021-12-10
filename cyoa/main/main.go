package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

const INITIATION_STRING = "intro"
const TERMINATION_STRING = "home"

func main() {
	jsonFilePointer := flag.String("file", "gopher.json", "The json file to load the story from")
	flag.Parse()

	jsonFile, err := ioutil.ReadFile(*jsonFilePointer)

	if err != nil {
		panic(err)
	}
	var story Story
	json.Unmarshal(jsonFile, &story)

	currentChapter := INITIATION_STRING

	tmpl := template.Must(template.ParseFiles("templates/template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, story[currentChapter])
			return
		}

		currentChapter = r.FormValue("arc")

		tmpl.Execute(w, story[currentChapter])
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)

}
