package model

type Product struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
