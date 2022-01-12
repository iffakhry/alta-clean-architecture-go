package main

import (
	_newsUsecase "alta/cleanarch/businesses/news"
	_newsController "alta/cleanarch/controllers/news"
	_driverFactory "alta/cleanarch/drivers"
	_newsRepo "alta/cleanarch/drivers/databases/news"
	"log"
	"time"

	_routes "alta/cleanarch/app/routes"

	_dbDriver "alta/cleanarch/drivers/mysql"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_newsRepo.News{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	newsRepo := _driverFactory.NewNewsRepository(db)
	newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, timeoutContext)
	newsCtrl := _newsController.NewNewsController(newsUsecase)

	routesInit := _routes.ControllerList{
		NewsController: *newsCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
