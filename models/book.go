package models

// Book model
type Book struct {
	Id       int     `json:"id" gorm:"primary_key"`
	Title    string  `json:"title" gorm:"unique"`
	Price    float64 `json:"price" gorm:"not null"`
	AuthorID int     `json:"author_id" gorm:"not null"`
	// Relationship
	Genres []Genre `gorm:"many2many:book_genres;"`
	Author *Author `gorm:"foreignKey:AuthorID"`
}
