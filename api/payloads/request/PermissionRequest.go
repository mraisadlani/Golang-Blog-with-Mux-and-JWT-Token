package request

import "errors"

type PermissionRequest struct {
	ID int64 `json:"id_permission" gorm:"column:id_permission"`
	IdMenu int64 `json:"id_menu" gorm:"column:id_menu"`
	IdUser int64 `json:"id_user" gorm:"column:id_user"`
	FCreate int64 `json:"f_create"`
	FRead int64 `json:"f_read"`
	FUpdate int64 `json:"f_update"`
	FDelete int64 `json:"f_delete"`
	FPublish int64 `json:"f_publish"`
	CreateAt string `json:"create_at"`
}

func (e *PermissionRequest) TableName() string {
	return "tb_has_permission"
}

func (e *PermissionRequest) Validate() error {
	if e.IdMenu == 0 {
		return errors.New("Id menu tidak ditemukan")
	}

	if e.IdUser == 0 {
		return errors.New("Id user tidak ditemukan")
	}

	return nil
}