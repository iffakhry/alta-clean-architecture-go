package news

import (
	"context"
	"time"
)

//buat struct core. tanpa ada gorm. karena kita ingin mengisolate domain business kita tanpa campur tangan library.
// gorm nanti dapat ditambahkan di repository.

type Domain struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//sebagai penyambung dari controller ke domain
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// GetByID(ctx context.Context, newsId int) (Domain, error)
	// GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	// Store(ctx context.Context, ip string, newsDomain *Domain) (Domain, error)
	// Update(ctx context.Context, newsDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// GetByID(ctx context.Context, newsId int) (Domain, error)
	// GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	// Store(ctx context.Context, newsDomain *Domain) (Domain, error)
	// Update(ctx context.Context, newsDomain *Domain) (Domain, error)
}
