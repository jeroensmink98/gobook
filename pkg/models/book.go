package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `gorm:"not null; unique; default:null" json:"title"`
	Isbn  string `gorm:"not null; unique; default:null" json:"isbn"`
}
