package models

type Article struct {
	Id              int
	Title           string
	CategaryId      int
	ArticleCategary ArticleCategary `gorm:"foreignKey:CategaryId;references:Id"`
	Tags            []Tags          `gorm:"many2many:article_tags;foreignKey:Id;joinForeignKey:ArticleID;References:Id;joinReferences:TagId"`
}

func (Article) TableName() string {
	return "article"
}
