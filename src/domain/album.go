package domain

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}

func NewAlbum(id string, title string, artist string, price int) (*Album, error) {
	return &Album{
		ID:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
	}, nil
}

func (a *Album) Update(title string, artist string, price int) error {
	a.Title = title
	a.Artist = artist
	a.Price = price
	return nil
}
