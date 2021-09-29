package entities

type PostCategory struct {
	ID int64 `json:"id" gorm:"column:id"`
	IdPost int64 `json:"id_post" gorm:"column:id_post"`
	IdCategory int64 `json:"id_category" gorm:"column:id_category"`
	Post *Post `gorm:"foreignKey:IdPost;references:ID;"`
	Category *Category `gorm:"foreignKey:IdCategory;references:ID"`
}

func (e *PostCategory) TableName() string {
	return "tb_posts_has_categories"
}