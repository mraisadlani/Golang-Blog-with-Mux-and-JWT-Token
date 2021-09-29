package entities

type Category struct {
	ID int64 `json:"id_category" gorm:"column:id_category"`
	NamaCategory string `json:"nama_category"`
	Slug string `json:"slug"`
	PostCategory *[]PostCategory `json:"posts" gorm:"foreignkey:IdCategory"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *Category) TableName() string {
	return "tb_categories"
}