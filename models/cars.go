package models

import "gorm.io/gorm"

type Cars struct {
	CarID       uint    `gorm:"primary key;autoIncrement" json:"carID"`
	CAR_NAME    *string `json:"carName"`
	CAR_HP      uint    `json:"carHp"`
	CAR_COMPANY *string `json:"carCompany"`
	CAR_ENGINE  *string `json:"carEngine"`
}

func MigrateCars(db *gorm.DB) error {
	err := db.AutoMigrate(&Cars{})
	return err
}
