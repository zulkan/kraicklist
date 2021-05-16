package domain

type Record struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	ThumbURL  string   `json:"thumb_url"`
	Tags      []string `json:"tags"`
	UpdatedAt int64    `json:"updated_at"`
	ImageURLs []string `json:"image_urls"`
}

type RecordRepository interface {
	FindAll()([]Record, error)
	FindAllTitle()([]string)
	FindByTitle(title string)(*Record, error) //return error if not found
	Count() int
}