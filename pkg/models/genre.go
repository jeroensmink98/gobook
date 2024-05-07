package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Genre string `gorm:"not null; unique; default:null" json:"genre"`
}
