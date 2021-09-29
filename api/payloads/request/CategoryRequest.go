package request

import (
	"errors"
	"github.com/gosimple/slug"
	"time"
)

type CategoryRequest struct {
	ID int64 `json:"id_category" gorm:"column:id_category"`
	NamaCategory string `json:"nama_category"`
	Slug string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *CategoryRequest) TableName() string {
	return "tb_categories"
}

func (e *CategoryRequest) Prepare() {
	e.Slug = slug.Make(e.NamaCategory)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *CategoryRequest) Validate() error {
	if e.NamaCategory == "" {
		return errors.New("Nama category tidak boleh kosong")
	}

	return nil
}