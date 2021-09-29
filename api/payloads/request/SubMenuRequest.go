package request

import (
	"errors"
	"github.com/gosimple/slug"
	"time"
)

type SubMenuRequest struct {
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

func (e *SubMenuRequest) TableName() string {
	return "tb_sub_menus"
}

func (e *SubMenuRequest) Prepare() {
	e.Slug = slug.Make(e.NamaSubMenu)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *SubMenuRequest) Validate() error {
	if e.NamaSubMenu == "" {
		return errors.New("Nama Sub Menu tidak boleh kosong")
	}

	if e.Icon == "" {
		return errors.New("Icon tidak boleh kosong")
	}

	if e.Path == "" {
		return errors.New("Path tidak boleh kosong")
	}

	return nil
}