package model

type Tag struct {
	Id    int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:Id"`
	Name  string `gorm:"column:Name"`
	Value string `gorm:"column:Value"`
}
