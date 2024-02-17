package models

type ArticleTags struct {
	ArticleId int
	TagId     int
}

func (ArticleTags) TableName() string {
	return "article_tags"
}
