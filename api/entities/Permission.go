package entities

type Permission struct {
	ID int64 `json:"id_permission" gorm:"column:id_permission"`
	IdMenu int64 `json:"id_menu" gorm:"column:id_menu"`
	IdUser int64 `json:"id_user" gorm:"column:id_user"`
	FCreate bool `json:"f_create"`
	FRead bool `json:"f_read"`
	FUpdate bool `json:"f_update"`
	FDelete bool `json:"f_delete"`
	FPublish bool `json:"f_publish"`
	CreateAt string `json:"create_at"`
	Menu *Menu `gorm:"foreignKey:IdMenu;references:ID;"`
	User *User `gorm:"foreignKey:IdUser;references:ID"`
}

func (e *Permission) TableName() string {
	return "tb_has_permission"
}