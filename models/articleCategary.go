package models

type ArticleCategary struct {
	Id       int
	Categary string
	Article  []Article `gorm:"foreignKey:CategaryId;references:Id"`
}

func (ArticleCategary) TableName() string {
	return "article_categary"
}
