package model

type Book struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Author    string   `json:"author"`
	CoAuthors []string `json:"co_authors"`
}
