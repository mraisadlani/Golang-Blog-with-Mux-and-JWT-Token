package entities

type Post struct {
	ID int64 `json:"id_post" gorm:"column:id_post"`
	NamaPost string `json:"nama_post"`
	Slug string `json:"slug"`
	Image string `json:"image"`
	Description string `json:"description"`
	Published bool `json:"published"`
	IdArticle int64 `json:"id_article" gorm:"column:id_article"`
	CreateBy string `json:"create_by"`
	CreateAt string `json:"create_at"`
	UpdateBy string `json:"update_by"`
	UpdateAt string `json:"update_at"`
	Category *[]PostCategory `gorm:"foreignKey:IdPost;references:ID"`
	Tag *[]PostTag `gorm:"foreignKey:IdPost;references:ID;"`
}

func (e *Post) TableName() string {
	return "tb_posts"
}