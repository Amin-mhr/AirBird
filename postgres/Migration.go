package postgres

import (
	"fmt"
	"gorm.io/gorm"
)

type AirPlane struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null"`
	AirPlaneType int    `gorm:"type:int;not null;foreignKey:AirPlaneTypeId"`
	SeatCount    int
	VipSeat      int    `gorm:"type:int;nullable"`
	Color_Out    string `gorm:"type:varchar(255);nullable"`
	Color_in     string `gorm:"type:varchar(255);nullable"`
	Cockpit_Type string `gorm:"type:varchar(255);nullable;ForeignKey:Cockpit_TypeId"`
}

type AirPlaneType struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
}

type Account struct {
	gorm.Model
	Username    string `gorm:"type:varchar(20);not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	AccountType int    `gorm:"type:int;not null;ForeignKey:AirPlaneTypeId"`
}

type Order struct {
	gorm.Model
	AirPlane  int `gorm:"type:int;not null;ForeignKey:AirPlaneId"`
	Account   int `gorm:"type:int;not null;ForeignKey:AccountId"`
	Situation int `gorm:"type:int;not null;ForeignKey:SituationId"`
}

type Situation struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
}

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type CockPit struct {
	gorm.Model
	model string `gorm:"type:varchar(20);not null"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&AirPlane{}, &AirPlaneType{}, &Account{}, &Order{}, &Situation{}, &Admin{})
	if err != nil {
		fmt.Println(err)
	}
}
