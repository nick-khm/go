package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:10])
}

func (p Post) HeadlineNoRef() string {
	p.Title = "Ruby"
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:10])
}

func (p *Post) Published() bool {
	return p.published
}

func (p *Post) Publish() {
	p.published = true
}

func (p *Post) Unpublish() {
	p.published = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {
	p := Post{
		Title: "Go release",
		Text:  "lorem100lorem100lorem100lorem100lorem100lorem100lorem100lorem100",
	}
	fmt.Println(p.Headline())
	fmt.Println(p.Published())
	p.Publish()
	fmt.Println(p.Published())
	p.Unpublish()
	fmt.Println(p.Published())
	UpperTitle(&p)
	fmt.Println(p.Headline())

	pythonPost := &Post{
		Title: "Python",
		Text:  "python python python python python python python python python python python python python python python ",
	}

	UpperTitle(pythonPost)
	fmt.Println(pythonPost.HeadlineNoRef())
	fmt.Println(pythonPost)

}
