package models

// title, content, author, date.
type Service struct{
	Title string `db:"title" json:"title"  form:"title"`
	Description string `db:"description" json:"decription" form:"decription"`
	Price float32 `db:"price" json:"price" form:"price"`
}
