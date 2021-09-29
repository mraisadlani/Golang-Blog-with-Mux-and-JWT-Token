package entities

type PostTag struct {
	ID int64 `json:"id" gorm:"column:id"`
	IdPost int64 `json:"id_post" gorm:"column:id_post"`
	IdTag int64 `json:"id_tag" gorm:"column:id_tag"`
	Post *Post `gorm:"foreignKey:IdPost;references:ID;"`
	Tag *Tag `gorm:"foreignKey:IdTag;references:ID"`
}

func (e *PostTag) TableName() string {
	return "tb_posts_has_tags"
}