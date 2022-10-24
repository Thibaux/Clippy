package Models

type Item struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      Tags   `json:"tags"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type Tag struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	TimesUsed int    `json:"times-used"`
}

type Items []Item

type Tags []Tag