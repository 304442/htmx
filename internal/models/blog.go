package models

// title, content, author, date.
type Blog struct{
	Title string `db:"title" json:"title"  form:"title"`
	Content string `db:"content" json:"content" form:"content"`
	Author string `db:"author" json:"author" form:"author"`
}