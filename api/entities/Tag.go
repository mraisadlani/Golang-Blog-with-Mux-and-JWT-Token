package entities

type Tag struct {
	ID int64 `json:"id_tag" gorm:"column:id_tag"`
	NamaTag string `json:"nama_tag"`
	Slug string `json:"slug"`
	PostTag *[]PostTag `json:"posts" gorm:"foreignkey:IdTag"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *Tag) TableName() string {
	return "tb_tags"
}