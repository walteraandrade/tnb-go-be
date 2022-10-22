// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comementary struct {
	ID     string  `json:"id"`
	Title  *string `json:"title"`
	Text   *string `json:"text"`
	Author *string `json:"author"`
}

type NewCommentary struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type NewPost struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Tags  []*string `json:"tags"`
}

type Post struct {
	ID    string    `json:"id"`
	Text  string    `json:"text"`
	Title string    `json:"title"`
	Tags  []*string `json:"tags"`
}
