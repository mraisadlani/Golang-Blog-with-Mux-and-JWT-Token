package entities

type SubMenu struct {
	ID int64 `json:"id_sub_menu" gorm:"column:id_sub_menu"`
	NamaSubMenu string `json:"nama_sub_menu"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
	Path string `json:"path"`
	Status bool `json:"status"`
	IdMenu int64 `json:"id_menu" gorm:"column:id_menu"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *SubMenu) TableName() string {
	return "tb_sub_menus"
}