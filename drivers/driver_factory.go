package drivers

import (
	newsDomain "alta/cleanarch/businesses/news"
	newsDB "alta/cleanarch/drivers/databases/news"

	"gorm.io/gorm"
)

//NewNewsRepository Factory with news domain
func NewNewsRepository(conn *gorm.DB) newsDomain.Repository {
	return newsDB.NewMySQLRepository(conn)
}
