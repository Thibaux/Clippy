package Models

type Item struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	UpdatedAt   string `json:"updatedAt"`
	CreatedAt   string `json:"createdAt"`
}

type Items []Item
