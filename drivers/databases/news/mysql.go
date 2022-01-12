package news

import (
	"alta/cleanarch/businesses/news"
	"context"

	"gorm.io/gorm"
)

type mysqlNewsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) news.Repository {
	return &mysqlNewsRepository{
		Conn: conn,
	}
}

func (repository *mysqlNewsRepository) GetAll(ctx context.Context) ([]news.Domain, error) {
	var data []News
	result := repository.Conn.Find(&data)
	if result.Error != nil {
		return []news.Domain{}, result.Error
	}

	var domainNews []news.Domain
	for _, value := range data {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, nil
}
