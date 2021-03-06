package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"car-prices/spiders"
	"log"
)

var (
	DB *gorm.DB

	username 	string = "kyo"
	password 	string = "kyo"
	dbName		string = "spiders"
)

func init()  {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf("gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		log.Println("default table name"+defaultTableName)
		return "sp_" + defaultTableName
	}
}

func AddCars(cars []spiders.QcCar)  {
	for index, car := range cars {
		log.Printf("%+v", car)
		if err := DB.Create(&car).Error; err != nil {
			log.Printf("db.Create index: %s, err: %v", index, err)
		}
	}
}
