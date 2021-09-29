package entities

type Article struct {
	ID int64 `json:"id_article" gorm:"column:id_article"`
	NamaArticle string `json:"nama_article"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
	Post *[]Post `json:"posts" gorm:"foreignKey:IdArticle;references:ID;"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *Article) TableName() string {
	return "tb_articles"
}