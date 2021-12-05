package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
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

	for currentChapter != TERMINATION_STRING {
		currentChapter = loadChapter(story[currentChapter])
	}

	loadChapter(story[TERMINATION_STRING])
}

func loadChapter(chapter Chapter) (nextChapter string) {
	color.HiGreen(chapter.Title)
	fmt.Printf("\n")

	for _, paragraph := range chapter.Paragraphs {
		color.HiYellow(paragraph)
		fmt.Scanln()
	}

	if len(chapter.Options) == 0 {
		return
	}

	c := color.New(color.FgCyan).Add(color.Underline)

	for idx, option := range chapter.Options {
		c.Println("Option", idx+1, "->", option.Text)
	}

	fmt.Printf("\n")

	var input int
	fmt.Scanln(&input)

	fmt.Printf("\n")

	nextChapter = chapter.Options[input-1].Chapter
	return nextChapter

}
