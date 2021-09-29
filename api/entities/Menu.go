package entities

type Menu struct {
	ID int64 `json:"id_menu" gorm:"column:id_menu"`
	NamaMenu string `json:"nama_menu"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
	Path string `json:"path"`
	Status bool `json:"status"`
	Permission *[]Permission `json:"permissions" gorm:"foreignkey:IdMenu"`
	SubMenu *[]SubMenu `json:"sub_menus" gorm:"foreignKey:IdMenu;references:ID;"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *Menu) TableName() string {
	return "tb_menus"
}