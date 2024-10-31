package server

type Article struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"size:100"`
	Content string `gorm:"unique;size:1000"`
}
