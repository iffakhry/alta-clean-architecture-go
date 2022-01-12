package request

import (
	"alta/cleanarch/businesses/news"
)

type News struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (req *News) ToDomain() *news.Domain {
	return &news.Domain{
		Title:   req.Title,
		Content: req.Content,
	}
}
