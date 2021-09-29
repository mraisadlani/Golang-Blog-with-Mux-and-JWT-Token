package entities

type Role struct {
	ID int64 `json:"id_role" gorm:"column:id_role"`
	NamaRole string `json:"nama_role"`
	Status bool `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *Role) TableName() string {
	return "tb_role"
}