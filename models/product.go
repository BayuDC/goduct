package models

type Product struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Desc string `gorm:"type:text" json:"description"`
}
