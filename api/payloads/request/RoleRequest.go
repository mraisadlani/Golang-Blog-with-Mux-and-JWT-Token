package request

import (
	"errors"
	"time"
)

type RoleRequest struct {
	ID int64 `json:"id_role" gorm:"column:id_role"`
	NamaRole string `json:"nama_role"`
	Status bool `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *RoleRequest) Prepare() {
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *RoleRequest) TableName() string {
	return "tb_role"
}

func (e *RoleRequest) Validate() error {
	if e.NamaRole == "" {
		return errors.New("Nama role tidak boleh kosong")
	}

	return nil
}