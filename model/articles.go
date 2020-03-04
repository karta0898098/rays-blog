package model

import "time"

type Article struct {
	Id          int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:Id"`
	Title       string     `gorm:"column:Title"`
	ReleaseDate *time.Time `gorm:"column:ReleaseDate"`
	Content     string     `gorm:"column:Content"`
	Tag         string     `gorm:"column:Tag"`
}
