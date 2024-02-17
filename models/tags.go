package models

type Tags struct {
	Id      int
	Tag     string
	Article []Article `gorm:"many2many:article_tags;foreignKey:Id;joinForeignKey:TagID;References:Id;joinReferences:ArticleId"`
}

func (Tags) TableName() string {
	return "tags"
}
