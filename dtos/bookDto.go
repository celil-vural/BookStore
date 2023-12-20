package dtos

type CreateBookDto struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	AuthorID int     `json:"author_id"`
	GenreIds []uint  `json:"genre_ids"`
}
